package service

import (
	"context"
	"github.com/MicroOps-cn/fuck-web/config"
	"math/rand"
	"runtime"
	"sync"

	"github.com/go-kit/log/level"
	"github.com/robfig/cron"

	logs "github.com/MicroOps-cn/fuck/log"
)

type localJobSchedule struct {
	mux         sync.Mutex
	concurrency int
	cron        *cron.Cron
	svc         JobService
}

func (s *localJobSchedule) OnJobHandlerChange(ctx context.Context) {
	jobCron, err := NewJobCron(ctx, config.Get().Job.Cron, s.svc)
	if err != nil {
		return
	}
	s.cron.Stop()
	jobCron.Start()
	s.cron = jobCron
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

func NewCronFunc(ctx context.Context, svc JobService, cronConfig *config.JobCron, jobId int, handler func(ctx context.Context, params string) error) func() {
	return func() {
		traceId := logs.NewTraceId()
		taskCtx, taskLogger := logs.NewContextLogger(ctx, logs.WithTraceId(traceId))
		err := svc.RunJob(taskCtx, cronConfig.Name, int32(jobId), int32(rand.Uint32()), traceId, func(ctx context.Context) error {
			if cronConfig.Timeout != nil && *cronConfig.Timeout != 0 {
				var cancelFunc context.CancelFunc
				ctx, cancelFunc = context.WithTimeout(ctx, *cronConfig.Timeout)
				defer cancelFunc()
			}
			return handler(ctx, cronConfig.Params)
		})
		if err != nil {
			level.Error(taskLogger).Log("err", err, "msg", "failed to run job", "name", cronConfig.Name, "expr", cronConfig.Expr)
			return
		}
	}
}

func NewJobCron(ctx context.Context, cronConfigs []*config.JobCron, svc JobService) (*cron.Cron, error) {
	cr := cron.New()
	for idx, cronConfig := range cronConfigs {
		handler := svc.GetJobHandler(ctx, cronConfig.Name)
		if handler == nil {
			continue
			//return nil, fmt.Errorf("the job `%s` not exists.", cronConfig.Name)
		}

		logger := logs.GetContextLogger(ctx)
		level.Info(logger).Log("msg", "add job to cron", "name", cronConfig.Name, "expr", cronConfig.Expr)
		if err := cr.AddFunc(cronConfig.Expr, NewCronFunc(ctx, svc, cronConfig, idx, handler)); err != nil {
			cr.Stop()
			return nil, err
		}
	}
	return cr, nil
}

func NewLocalJobSchedule(ctx context.Context, jobSvc JobService) (*localJobSchedule, error) {
	jobCron, err := NewJobCron(ctx, config.Get().Job.Cron, jobSvc)
	if err != nil {
		return nil, err
	}
	localSchedule := &localJobSchedule{concurrency: runtime.NumCPU(), cron: jobCron, svc: jobSvc}

	config.OnConfigReload(func(o, n *config.Config) error {
		jobCron, err = NewJobCron(ctx, config.Get().Job.Cron, jobSvc)
		if err != nil {
			return err
		}
		localSchedule.cron.Stop()
		jobCron.Start()
		localSchedule.cron = jobCron
		return nil
	}, func(o, n *config.Config) {})
	jobCron.Start()
	return localSchedule, nil
}
