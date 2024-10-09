package service

import (
	"context"
	"fmt"
	"time"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/service/gormservice"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
)

type LoggingService interface {
	migrator
	PostEventLog(ctx context.Context, eventId, userId, username, clientIP, loc, action, message string, status bool, took time.Duration, log ...interface{}) error
	GetEvents(ctx context.Context, filters map[string]string, keywords string, startTime time.Time, endTime time.Time, current int64, size int64) (count int64, event []*models.Event, err error)
	GetEventLogs(ctx context.Context, filters map[string]string, keywords string, current int64, size int64) (count int64, event []*models.EventLog, err error)
}

func NewLoggingService(ctx context.Context) LoggingService {
	var loggingService LoggingService
	loggingStorage := config.Get().GetStorage().GetLogging()
	switch loggingSource := loggingStorage.GetStorageSource().(type) {
	case *config.Storage_Mysql:
		loggingService = gormservice.NewLoggingService(ctx, loggingStorage.Name, loggingSource.Mysql.Client)
	case *config.Storage_Sqlite:
		loggingService = gormservice.NewLoggingService(ctx, loggingStorage.Name, loggingSource.Sqlite.Client)
	default:
		panic(fmt.Sprintf("failed to initialize LoggingService: unknown data source: %T", loggingSource))
	}
	return loggingService
}
