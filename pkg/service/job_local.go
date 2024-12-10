package service

import (
	"context"
	"sync"

	"github.com/go-kit/log/level"
	"github.com/robfig/cron"

	logs "github.com/MicroOps-cn/fuck/log"
)

type localJobSchedule struct {
	mux         sync.Mutex
	concurrency int
	cron        *cron.Cron
}

func (s *localJobSchedule) AddFunc(spec string, cmd func()) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.cron.AddFunc(spec, cmd)
}

func (s *localJobSchedule) GetConcurrency() int {
	return s.concurrency
}

func (s *localJobSchedule) Callback(ctx context.Context, id int32, code int, msg string) error {
	logger := logs.GetContextLogger(ctx)
	if code != 200 {
		level.Error(logger).Log("err", msg, "msg", "failed to run task", "id", id)
	} else {
		level.Debug(logger).Log("msg", "task finish", "id", id)
	}
	return nil
}

func (s *localJobSchedule) Reset() {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.cron != nil {
		s.cron.Stop()
	}
	s.cron = cron.New()
	s.cron.Start()
}
