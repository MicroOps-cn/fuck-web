package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/MicroOps-cn/fuck/clients/gorm"
	"github.com/MicroOps-cn/fuck/conv"
	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/sets"
	"github.com/go-kit/log/level"
	uuid "github.com/satori/go.uuid"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/service/gormservice"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
)

type CommonService interface {
	baseService
	RecordUploadFile(ctx context.Context, name string, path string, contentType string, size int64) (id string, err error)
	GetFileInfoFromId(ctx context.Context, id string) (fileName, mimiType, filePath string, err error)
	CreateRole(ctx context.Context, role *models.Role) (err error)
	UpdateRole(ctx context.Context, role *models.Role) (err error)
	GetRoles(ctx context.Context, keywords string, current, pageSize int64) (count int64, roles []*models.Role, err error)
	GetPermissions(ctx context.Context, keywords string, current int64, pageSize int64) (count int64, permissions []*models.Permission, err error)
	DeleteRoles(ctx context.Context, ids []string) error
	RegisterPermission(ctx context.Context, permissions models.Permissions) error
	CreateOrUpdateRoleByName(ctx context.Context, role *models.Role) error
	Authorization(ctx context.Context, roles []string, method string) bool

	GetUserExtendedData(ctx context.Context, id string) (*models.UserExt, error)
	GetUsersExtendedData(ctx context.Context, id []string) ([]*models.UserExt, error)
	PatchUserExtData(ctx context.Context, id string, patch map[string]interface{}) error
	PatchUserRole(ctx context.Context, id string, roleName, roleId string) error
	GetUserRole(ctx context.Context, id string) (role *models.Role, err error)

	CreateTOTP(ctx context.Context, ids string, secret string) error
	GetTOTPSecrets(ctx context.Context, ids []string) ([]string, error)
	PatchSystemConfig(ctx context.Context, prefix string, patch map[string]interface{}) error
	GetSystemConfig(ctx context.Context, prefix string) (map[string]interface{}, error)
	VerifyAndRecordHistoryPassword(ctx context.Context, id string, password string) error
	UpdateLoginTime(ctx context.Context, id string) error
	VerifyWeakPassword(ctx context.Context, password string) error
	InsertWeakPassword(ctx context.Context, passwords ...string) error
	GetUsersRole(ctx context.Context, ids []string) ([]*models.UserRole, error)
}

func NewCommonService(ctx context.Context) CommonService {
	// logger := log.With(logs.GetContextLogger(ctx), "service", "common")
	// ctx = context.WithValue(ctx, global.LoggerName, logger)
	var commonService CommonService
	commonStorage := config.Get().GetStorage().GetDefault()
	switch commonSource := commonStorage.GetStorageSource().(type) {
	case *config.Storage_Mysql:
		commonService = gormservice.NewCommonService(commonStorage.Name, commonSource.Mysql.Client)
	case *config.Storage_Sqlite:
		commonService = gormservice.NewCommonService(commonStorage.Name, commonSource.Sqlite.Client)
	default:
		panic(fmt.Sprintf("failed to initialize CommonService: unknown data source: %T", commonSource))
	}
	return commonService
}

func (s Set) LoadSystemConfig(ctx context.Context) error {
	cfgs, err := s.commonService.GetSystemConfig(ctx, "security")
	if err != nil {
		return fmt.Errorf("failed to load runtime config: %s", err)
	}
	config.SetRuntimeConfig(func(c *config.RuntimeConfig) {
		if err = conv.JSON(cfgs, c.Security); err != nil {
			err = fmt.Errorf("failed to parse runtime config: %s", err)
		}
	})
	return err
}

func (s Set) UploadFile(ctx context.Context, name, contentType string, f io.Reader) (fileKey string, err error) {
	logger := logs.GetContextLogger(ctx)
	now := time.Now().UTC()
	d, err := config.Get().GetUploadDir()
	if err != nil {
		level.Error(logger).Log("err", err, "msg", "failed to get upload dir")
		return "", errors.InternalServerError()
	}
	dirName := now.Format("2006-01")
	if _, err = d.Stat(dirName); os.IsNotExist(err) {
		if err = d.MkdirAll(dirName, 0o755); err != nil {
			level.Error(logger).Log("msg", "failed to create directory", "err", err)
		}
	} else if err != nil {
		level.Error(logger).Log("msg", "failed to get directory status", "err", err)
	}
	filePath := fmt.Sprintf("%s/%d%s", dirName, now.UnixNano(), path.Ext(name))

	var ff io.ReadWriteCloser
	if ff, err = d.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644); err != nil {
		level.Error(logger).Log("err", err, "msg", "failed to open file", "filePath", filePath)
		return "", errors.InternalServerError()
	}
	defer ff.Close()
	size, err := io.Copy(ff, f)
	if err != nil {
		return "", err
	}
	return s.commonService.RecordUploadFile(ctx, name, filePath, contentType, size)
}

func (s Set) DownloadFile(ctx context.Context, id string) (f io.ReadCloser, mimiType, fileName string, err error) {
	var filePath string
	fileName, mimiType, filePath, err = s.commonService.GetFileInfoFromId(ctx, id)
	if err != nil {
		return nil, "", "", err
	}
	logger := logs.GetContextLogger(ctx)
	d, err := config.Get().GetUploadDir()
	if err != nil {
		level.Error(logger).Log("err", err, "msg", "failed to get upload dir")
		return nil, "", "", errors.InternalServerError()
	}
	if f, err = d.Open(filePath); err != nil {
		level.Error(logger).Log("err", err, "msg", "failed to open file", "filePath", filePath)
		return nil, "", "", errors.NotFoundError()
	}
	return f, mimiType, fileName, nil
}

func (s Set) CreateRole(ctx context.Context, role *models.Role) (err error) {
	return s.commonService.CreateRole(ctx, role)
}

func (s Set) UpdateRole(ctx context.Context, role *models.Role) (err error) {
	return s.commonService.UpdateRole(ctx, role)
}

func (s Set) GetRoles(ctx context.Context, keywords string, current, pageSize int64) (count int64, roles []*models.Role, err error) {
	return s.commonService.GetRoles(ctx, keywords, current, pageSize)
}

func (s Set) GetPermissions(ctx context.Context, keywords string, current int64, pageSize int64) (count int64, permissions []*models.Permission, err error) {
	return s.commonService.GetPermissions(ctx, keywords, current, pageSize)
}

func (s Set) DeleteRoles(ctx context.Context, ids []string) error {
	return s.commonService.DeleteRoles(ctx, ids)
}

func (s Set) Authorization(ctx context.Context, user *models.User, method string) bool {
	roles := sets.New[string](user.Role)
	return s.commonService.Authorization(ctx, roles.List(), method)
}

func (s Set) RegisterPermission(ctx context.Context, permissions models.Permissions) error {
	err := s.commonService.RegisterPermission(ctx, permissions)
	if err != nil {
		return err
	}

	for _, role := range permissions.GetRoles() {
		if err = s.commonService.CreateOrUpdateRoleByName(ctx, role); err != nil {
			return err
		}
	}
	return nil
}

func GetEventMeta(ctx context.Context, action string, beginTime time.Time, err error) (eventId, message string, status bool, took time.Duration) {
	eventId = logs.GetTraceId(ctx)
	if u, e := uuid.FromString(eventId); e == nil {
		eventId = u.String()
	}
	if err != nil {
		return eventId, fmt.Sprintf("Calling the %s method failed, err: %s", action, err), false, time.Since(beginTime)
	}
	return eventId, fmt.Sprintf("Successfully called %s method.", action), true, time.Since(beginTime)
}

func (s Set) InsertWeakPassword(ctx context.Context, passwords ...string) error {
	return s.commonService.InsertWeakPassword(ctx, passwords...)
}

func (s Set) VerifyWeakPassword(ctx context.Context, password string) error {
	return s.commonService.VerifyWeakPassword(ctx, password)
}

func NewDefaultDB() *gorm.Client {
	commonStorage := config.Get().GetStorage().GetDefault()
	switch commonSource := commonStorage.GetStorageSource().(type) {
	case *config.Storage_Mysql:
		return commonSource.Mysql.Client
	case *config.Storage_Sqlite:
		return commonSource.Sqlite.Client
	default:
		panic(fmt.Sprintf("failed to initialize CommonService: unknown data source: %T", commonSource))
	}
	return nil
}
