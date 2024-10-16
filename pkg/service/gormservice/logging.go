package gormservice

import (
	"context"
	"fmt"
	"sync"
	"time"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/signals"
	"github.com/go-kit/log/level"
	uuid "github.com/satori/go.uuid"
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/httputil"
	"github.com/MicroOps-cn/fuck/clients/gorm"
)

func NewLoggingService(ctx context.Context, name string, client *gorm.Client) *LoggingService {
	svc := &LoggingService{name: name, Client: client, buf: EventBuffers{event: make(map[string]*models.Event)}}
	go svc.Pusher(ctx)
	return svc
}

type EventBuffers struct {
	event map[string]*models.Event
	logs  []*models.EventLog
	mux   sync.Mutex
}

func (b *EventBuffers) PutEvent(ctx context.Context, eventId, userId, username, clientIP, loc, action, message string, status bool, took time.Duration, logItems ...interface{}) error {
	if len(eventId) == 0 {
		if eId, e := uuid.FromString(logs.GetTraceId(ctx)); e != nil {
			eventId = logs.NewTraceId()
		} else {
			eventId = eId.String()
		}
	}

	b.mux.Lock()
	defer b.mux.Unlock()
	if event, ok := b.event[eventId]; ok {
		event.UserId = httputil.NewValue(userId, httputil.Default(event.UserId)).String()
		event.Username = httputil.NewValue(username, httputil.Default(event.Username)).String()
		event.ClientIP = httputil.NewValue(clientIP, httputil.Default(event.ClientIP)).String()
		event.Location = httputil.NewValue(loc, httputil.Default(event.Location)).String()
		event.Action = httputil.NewValue(action, httputil.Default(event.Action)).String()
		event.Message = httputil.NewValue(message, httputil.Default(event.Message)).String()
		event.Status = status
		event.Took = took
		event.CreateTime = time.Now().UTC()
	} else {
		b.event[eventId] = &models.Event{UserId: userId, Username: username, ClientIP: clientIP, Location: loc, Action: action, Message: message, Status: status, Took: took}
		b.event[eventId].Id = eventId
		b.event[eventId].CreateTime = time.Now().UTC()
	}

	var errs errors.MultipleServerError
	for _, item := range logItems {
		eventLog := models.EventLog{Model: models.Model{CreateTime: time.Now().UTC()}, EventId: eventId}
		switch v := item.(type) {
		case []byte:
			eventLog.Log = v
		case string:
			eventLog.Log = models.CompressField(v)
		default:
			data, err := json.Marshal(v)
			if err != nil {
				errs.Append(err)
				continue
			}
			eventLog.Log = data
		}
		b.logs = append(b.logs, &eventLog)
	}
	if errs.HasError() {
		return errs
	}
	return nil
}

func (b *EventBuffers) getEvents() ([]models.Event, []models.EventLog) {
	b.mux.Lock()
	defer b.mux.Unlock()
	var events []models.Event
	var log []models.EventLog
	if len(b.event) > 0 {
		newEvents := make(map[string]*models.Event)
		for _, event := range b.event {
			if time.Now().UTC().Sub(event.CreateTime) > time.Second*3 {
				events = append(events, *event)
			} else {
				newEvents[event.Id] = event
			}
		}
		b.event = newEvents
	}
	if len(b.logs) > 0 {
		for _, l := range b.logs {
			log = append(log, *l)
		}
		b.logs = []*models.EventLog{}
	}
	return events, log
}

func (s *LoggingService) Pusher(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 3)
	stopCh := signals.SetupSignalHandler(logs.GetContextLogger(ctx))
	stopCh.AddRequest(1)
loop:
	for {
		logger := logs.NewTraceLogger()
		conn := s.Session(ctx)
		select {
		case <-ticker.C:
			events, log := s.buf.getEvents()
			if len(events) > 0 {
				if err := conn.CreateInBatches(events, 100).Error; err != nil {
					level.Error(logger).Log("msg", "failed to push event", "err", err)
				}
			}
			if len(log) > 0 {
				if err := conn.CreateInBatches(log, 100).Error; err != nil {
					level.Error(logger).Log("msg", "failed to push event logs", "err", err)
				}
			}
		case <-ctx.Done():
			events, log := s.buf.getEvents()
			if len(events) > 0 {
				if err := conn.CreateInBatches(events, 100).Error; err != nil {
					level.Error(logger).Log("msg", "failed to push event", "err", err)
				}
			}
			if len(log) > 0 {
				if err := conn.CreateInBatches(log, 100).Error; err != nil {
					level.Error(logger).Log("msg", "failed to push event logs", "err", err)
				}
			}
			level.Debug(logs.GetContextLogger(ctx)).Log("msg", "close event pusher")
			break loop
		}
	}
	stopCh.DoneRequest()
}

type LoggingService struct {
	*gorm.Client
	name string
	buf  EventBuffers
}

func (s *LoggingService) Name() string {
	return s.name
}

func (s *LoggingService) AutoMigrate(ctx context.Context) error {
	return s.Session(ctx).AutoMigrate(
		&models.Event{},
		&models.EventLog{},
	)
}

func (s *LoggingService) PostEventLog(ctx context.Context, eventId, userId, username, clientIP, loc, action, message string, status bool, took time.Duration, logItems ...interface{}) error {
	return s.buf.PutEvent(ctx, eventId, userId, username, clientIP, loc, action, message, status, took, logItems...)
}

func (s *LoggingService) GetEvents(ctx context.Context, filters map[string]string, keywords string, startTime time.Time, endTime time.Time, current int64, pageSize int64) (count int64, events []*models.Event, err error) {
	conn := s.Session(ctx)
	tb := conn.Model(&models.Event{}).Where("create_time > ? and create_time < ?", startTime, endTime)
	if len(filters) > 0 {
		tb = tb.Where(filters)
	}
	if len(keywords) > 0 {
		keywords = fmt.Sprintf("%%%s%%", keywords)
		tb = tb.Where(conn.
			Where("user_id like ?", keywords).
			Or("username like ?", keywords).
			Or("action like ?", keywords).
			Or("client_ip like ?", keywords).
			Or("location like ?", keywords).
			Or("message like ?", keywords),
		)
	}

	if err = tb.Count(&count).Error; err != nil {
		return 0, nil, err
	} else if count == 0 {
		return 0, nil, nil
	}

	if err = tb.Order("`id` DESC").Limit(int(pageSize)).Offset(int((current - 1) * pageSize)).Find(&events).Error; err != nil {
		return 0, nil, err
	}
	return count, events, nil
}

func (s *LoggingService) GetEventLogs(ctx context.Context, filters map[string]string, keywords string, current int64, pageSize int64) (count int64, eventLogs []*models.EventLog, err error) {
	conn := s.Session(ctx)
	tb := conn.Model(&models.EventLog{}).Joins("JOIN `t_event` ON `t_event`.`id` = `t_event_log`.`event_id`")

	if len(filters) > 0 {
		tb = tb.Where(filters)
	}
	if len(keywords) > 0 {
		keywords = fmt.Sprintf("%%%s%%", keywords)
		tb = tb.Where(conn.
			Where("event_id like ?", keywords).
			Or("log like ?", keywords),
		)
	}

	if err = tb.Count(&count).Error; err != nil {
		return 0, nil, err
	} else if count == 0 {
		return 0, nil, nil
	}

	if err = tb.Limit(int(pageSize)).Offset(int((current - 1) * pageSize)).Select("t_event_log.*").Find(&eventLogs).Error; err != nil {
		return 0, nil, err
	}
	return count, eventLogs, nil
}
