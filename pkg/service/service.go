package service

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/sets"
	"github.com/go-kit/log/level"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/client/email"
	"github.com/MicroOps-cn/fuck-web/pkg/client/geoip"
	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/service/gormservice"
	"github.com/MicroOps-cn/fuck-web/pkg/service/ldapservice"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck-web/pkg/service/opts"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/sign"
)

type migrator interface {
	AutoMigrate(ctx context.Context) error
}

type baseService interface {
	migrator
}

type Service interface {
	baseService
	InitData(ctx context.Context, username string) error
	DeleteLoginSession(ctx context.Context, session string) error
	GetSessionByToken(ctx context.Context, id string, tokenType models.TokenType, receiver interface{}) error
	VerifyPassword(ctx context.Context, username string, password string, allowPasswordExpired bool) (user *models.User, err error)
	VerifyPasswordById(ctx context.Context, userId, password string, allowPasswordExpired bool) (user *models.User)
	VerifyUserStatus(ctx context.Context, user *models.User, allowPasswordExpired bool) (err error)

	GetSessions(ctx context.Context, userId string, current int64, size int64) (int64, []*models.Token, error)
	DeleteToken(ctx context.Context, tokenType models.TokenType, id string) (err error)

	UploadFile(ctx context.Context, name, contentType string, f io.Reader) (fileKey string, err error)
	DownloadFile(ctx context.Context, id string) (f io.ReadCloser, mimiType, fileName string, err error)

	CreateRole(ctx context.Context, role *models.Role) (err error)
	UpdateRole(ctx context.Context, role *models.Role) (err error)
	GetRoles(ctx context.Context, keywords string, current, pageSize int64) (count int64, roles []*models.Role, err error)
	GetPermissions(ctx context.Context, keywords string, current int64, pageSize int64) (count int64, permissions []*models.Permission, err error)
	DeleteRoles(ctx context.Context, ids []string) error

	GetUsers(ctx context.Context, keywords string, status models.UserMeta_UserStatus, current, pageSize int64) (total int64, users models.Users, err error)
	PatchUsers(ctx context.Context, patch []map[string]interface{}) (count int64, err error)
	DeleteUsers(ctx context.Context, id []string) (count int64, err error)
	UpdateUser(ctx context.Context, user *models.User, updateColumns ...string) (err error)
	GetUserInfo(ctx context.Context, id, username string) (user *models.User, err error)
	GetUser(ctx context.Context, options ...opts.WithGetUserOptions) (user *models.User, err error)
	GetUserInfoByUsernameAndEmail(ctx context.Context, username, email string) (user *models.User, err error)
	CreateUser(ctx context.Context, user *models.User) (err error)
	PatchUser(ctx context.Context, user map[string]interface{}) (err error)
	DeleteUser(ctx context.Context, id string) error
	PatchUserExtData(ctx context.Context, id string, m map[string]interface{}) error
	Authentication(ctx context.Context, method models.AuthMeta_Method, algorithm sign.AuthAlgorithm, key, secret, payload, signStr string) (*models.User, error)

	ResetPassword(ctx context.Context, id, password string) error

	UpdateToken(ctx context.Context, id string, tokenType models.TokenType, data interface{}) (err error)
	UpdateUserSession(ctx context.Context, userId string) (err error)
	CreateToken(ctx context.Context, tokenType models.TokenType, data interface{}) (token *models.Token, err error)
	VerifyToken(ctx context.Context, token string, tokenType models.TokenType, receiver interface{}, relationId ...string) bool
	SendEmail(ctx context.Context, data map[string]interface{}, topic string, to ...string) error
	Authorization(ctx context.Context, user *models.User, method string) bool
	RegisterPermission(ctx context.Context, permissions models.Permissions) error

	CreateTOTP(ctx context.Context, ids string, secret string) error
	GetTOTPSecrets(ctx context.Context, ids []string) ([]string, error)
	PatchSystemConfig(ctx context.Context, prefix string, patch map[string]interface{}) error
	LoadSystemConfig(ctx context.Context) error
	PostEventLog(ctx context.Context, eventId, userId, username, clientIP, action, message string, status bool, took time.Duration, log ...interface{}) error
	GetEvents(ctx context.Context, filters map[string]string, keywords string, startTime time.Time, endTime time.Time, current int64, size int64) (count int64, event []*models.Event, err error)
	GetEventLogs(ctx context.Context, filters map[string]string, keywords string, current int64, size int64) (count int64, event []*models.EventLog, err error)
	InsertWeakPassword(ctx context.Context, passwords ...string) error
	VerifyWeakPassword(ctx context.Context, password string) error
	UpdateTokenExpires(ctx context.Context, id string, expiry time.Time) error
	GetUserRole(ctx context.Context, id string) (*models.Role, error)
}

type Set struct {
	userService    UserService
	sessionService SessionService
	commonService  CommonService
	loggingService LoggingService
	geoIPClient    *geoip.Client
}

func (s Set) GetEvents(ctx context.Context, filters map[string]string, keywords string, startTime, endTime time.Time, current int64, size int64) (count int64, event []*models.Event, err error) {
	return s.loggingService.GetEvents(ctx, filters, keywords, startTime, endTime, current, size)
}

func (s Set) GetEventLogs(ctx context.Context, filters map[string]string, keywords string, current int64, size int64) (count int64, event []*models.EventLog, err error) {
	return s.loggingService.GetEventLogs(ctx, filters, keywords, current, size)
}

func (s Set) PostEventLog(ctx context.Context, eventId, userId, username, clientIP, action, message string, status bool, took time.Duration, log ...interface{}) (err error) {
	var loc string
	if s.geoIPClient != nil {
		logger := logs.GetContextLogger(ctx)
		loc, err = s.geoIPClient.City(net.ParseIP(clientIP))
		if err != nil {
			level.Error(logger).Log("msg", "failed to convert ip to location", "err", err, "clientIP", clientIP)
		}
	}
	return s.loggingService.PostEventLog(ctx, eventId, userId, username, clientIP, loc, action, message, status, took, log...)
}

func (s Set) GetUserService() UserService {
	return s.userService
}

func (s Set) SendEmail(ctx context.Context, data map[string]interface{}, topic string, to ...string) error {
	if len(to) == 0 {
		level.Error(logs.GetContextLogger(ctx)).Log("err")
		return errors.ParameterError("recipient is empty")
	}

	nowTs := time.Now().Unix()
	ts := nowTs - nowTs%60
	seed := fmt.Sprintf("%s|%s|%d", topic, strings.Join(sets.New[string](to...).SortedList(), ","), ts)
	count, err := s.sessionService.GetCounter(ctx, seed)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.NewServerError(429, "the sending frequency is too fast. please try again in 60 seconds", errors.CodeRequestTooFrequently)
	}
	expr := time.Now().UTC().Add(time.Minute)
	if err = s.sessionService.Counter(ctx, seed, &expr); err != nil {
		return err
	}
	smtpConfig := config.Get().Smtp
	if smtpConfig == nil {
		return errors.NewServerError(500, "failed to get smtp options")
	}
	subject, body, err := smtpConfig.GetSubjectAndBody(data, topic)
	if err != nil {
		return errors.WithServerError(500, err, fmt.Sprintf("failed to get email body: topic=%s ", topic))
	}
	client, err := email.NewSMTPClient(ctx, smtpConfig)
	if err != nil {
		return errors.WithServerError(500, err, "failed to create SMTP client")
	}
	client.SetSubject(subject)
	client.SetBody("text/html", body)
	client.SetTo(to)
	return client.Send()
}

func (s Set) CreateToken(ctx context.Context, tokenType models.TokenType, data interface{}) (token *models.Token, err error) {
	tk, err := models.NewToken(tokenType, data)
	if err != nil {
		return nil, err
	}
	err = s.sessionService.CreateToken(ctx, tk)
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (s Set) UpdateToken(ctx context.Context, id string, tokenType models.TokenType, data interface{}) (err error) {
	tk, err := models.NewToken(tokenType, data)
	if err != nil {
		return err
	}
	tk.Id = id
	err = s.sessionService.UpdateToken(ctx, tk)
	if err != nil {
		return err
	}
	return nil
}

func (s Set) UpdateTokenExpires(ctx context.Context, id string, expiry time.Time) (err error) {
	return s.sessionService.UpdateTokenExpires(ctx, id, expiry)
}

func (s Set) VerifyToken(ctx context.Context, token string, tokenType models.TokenType, receiver interface{}, relationId ...string) bool {
	logger := logs.GetContextLogger(ctx)
	if len(token) == 0 {
		return false
	}
	tk, err := s.sessionService.GetToken(ctx, token, tokenType, relationId...)

	if err != nil {
		return false
	} else if tk == nil {
		return false
	}
	switch tokenType {
	case models.TokenTypeCode,
		models.TokenTypeOAuthState,
		models.TokenTypeLoginCode:
		if err = s.DeleteToken(ctx, tokenType, tk.Id); err != nil {
			level.Warn(logger).Log("msg", "failed to delete token.", "err", err)
		}
	}
	if !tk.Expiry.After(time.Now().UTC()) {
		return false
	}
	if receiver != nil {
		if err = tk.To(receiver); err != nil {
			level.Warn(logger).Log("msg", "failed to parse token data.", "err", err)
			return false
		}
	}
	return true
}

func (s Set) InitData(ctx context.Context, username string) error {
	adminUser, err := s.GetUserInfo(ctx, "", username)
	if errors.IsNotFount(err) {
		adminUser = &models.User{
			Username: username,
			Password: sql.RawBytes("fuck-web"),
			Role:     "admin",
		}
		err = s.CreateUser(ctx, adminUser)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func (s Set) AutoMigrate(ctx context.Context) error {
	svcs := []baseService{
		s.commonService, s.sessionService, s.userService, s.loggingService,
	}
	for _, svc := range svcs {
		if err := svc.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(ctx context.Context) Service {
	return &Set{
		userService:    NewUserAndAppService(ctx),
		sessionService: NewSessionService(ctx),
		commonService:  NewCommonService(ctx),
		loggingService: NewLoggingService(ctx),
		geoIPClient:    config.Get().GetStorage().Geoip,
	}
}

type UserService interface {
	baseService

	Name() string
	GetUsers(ctx context.Context, keywords string, status models.UserMeta_UserStatus, current, pageSize int64) (total int64, users []*models.User, err error)
	PatchUsers(ctx context.Context, patch []map[string]interface{}) (count int64, err error)
	DeleteUsers(ctx context.Context, id []string) (count int64, err error)
	UpdateUser(ctx context.Context, user *models.User, updateColumns ...string) (err error)
	GetUserInfo(ctx context.Context, id string, username string) (*models.User, error)
	GetUser(ctx context.Context, options *opts.GetUserOptions) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (err error)
	PatchUser(ctx context.Context, user map[string]interface{}) (err error)
	DeleteUser(ctx context.Context, id string) error

	VerifyPassword(ctx context.Context, username string, password string) *models.User
	ResetPassword(ctx context.Context, id string, password string) error
	GetUserInfoByUsernameAndEmail(ctx context.Context, username, email string) (*models.User, error)
	VerifyPasswordById(ctx context.Context, id, password string) (user *models.User)
	GetUsersById(ctx context.Context, id []string) (models.Users, error)
}

type UserServices []UserService

func NewUserAndAppService(ctx context.Context) UserService {
	ctx, _ = logs.NewContextLogger(ctx, logs.WithKeyValues("service", "userAndApp"))
	if config.Get().GetStorage().GetUser() != nil {
		userStorage := config.Get().GetStorage().GetUser()
		switch userSource := userStorage.GetStorageSource().(type) {
		case *config.Storage_Mysql:
			return gormservice.NewUserAndAppService(ctx, userStorage.GetName(), userSource.Mysql.Client)
		case *config.Storage_Sqlite:
			return gormservice.NewUserAndAppService(ctx, userStorage.GetName(), userSource.Sqlite.Client)
		case *config.Storage_Ldap:
			return ldapservice.NewUserAndAppService(ctx, userStorage.GetName(), userSource.Ldap)
		default:
			panic(fmt.Sprintf("Failed to init UserService: Unknown datasource: %T ", userSource))
		}

	} else {
		panic("Failed to init UserService: user source is not set")
	}
}
