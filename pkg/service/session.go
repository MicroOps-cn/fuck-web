package service

import (
	"context"
	"fmt"
	"time"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/go-kit/log/level"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/service/gormservice"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck-web/pkg/service/redisservice"
)

type SessionService interface {
	baseService

	GetSessions(ctx context.Context, userId string, current int64, size int64) (int64, []*models.Token, error)
	DeleteToken(ctx context.Context, tokenType models.TokenType, id string) (err error)
	GetToken(ctx context.Context, token string, tokenType models.TokenType, relationId ...string) (*models.Token, error)
	CreateToken(ctx context.Context, token *models.Token) error
	UpdateToken(ctx context.Context, token *models.Token) error
	Counter(ctx context.Context, seed string, expireTime *time.Time, num ...int64) (err error)
	GetCounter(ctx context.Context, seed string) (count int64, err error)
	UpdateTokenExpires(ctx context.Context, id string, expiry time.Time) error
}

func NewSessionService(ctx context.Context) SessionService {
	// logger := log.With(logs.GetContextLogger(ctx), "service", "session")
	// ctx = context.WithValue(ctx, global.LoggerName, logger)
	var sessionService SessionService
	sessionStorage := config.Get().GetStorage().GetSession()
	switch sessionSource := sessionStorage.GetStorageSource().(type) {
	case *config.Storage_Mysql:
		sessionService = gormservice.NewSessionService(ctx, sessionStorage.Name, sessionSource.Mysql.Client)
	case *config.Storage_Sqlite:
		sessionService = gormservice.NewSessionService(ctx, sessionStorage.Name, sessionSource.Sqlite.Client)
	case *config.Storage_Redis:
		sessionService = redisservice.NewSessionService(sessionStorage.Name, sessionSource.Redis)
	default:
		panic(fmt.Sprintf("failed to initialize SessionService: unknown data source: %T", sessionSource))
	}
	return sessionService
}

func (s Set) DeleteLoginSession(ctx context.Context, sessionId string) error {
	return s.sessionService.DeleteToken(ctx, models.TokenTypeLoginSession, sessionId)
}

func (s Set) GetSessionByToken(ctx context.Context, id string, tokenType models.TokenType, receiver interface{}) (err error) {
	logger := logs.GetContextLogger(ctx)
	token, err := s.sessionService.GetToken(ctx, id, tokenType)
	if err != nil {
		return err
	}

	if time.Since(token.LastSeen) > time.Minute {
		token.LastSeen = time.Now().UTC()
		if err = s.sessionService.UpdateToken(ctx, token); err != nil {
			level.Warn(logger).Log("msg", "failed to update token last seen.", "err", err)
		}
	}
	return token.To(receiver)
}

func (s Set) GetSessions(ctx context.Context, userId string, current, size int64) (int64, []*models.Token, error) {
	return s.sessionService.GetSessions(ctx, userId, current, size)
}

func (s Set) DeleteToken(ctx context.Context, tokenType models.TokenType, id string) (err error) {
	return s.sessionService.DeleteToken(ctx, tokenType, id)
}
