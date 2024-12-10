package service

import (
	"bufio"
	"context"
	errors2 "errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"runtime"
	"slices"
	"sync"
	"time"

	"github.com/go-kit/kit/metrics/prometheus"
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/robfig/cron"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/client/xxljob"
	"github.com/MicroOps-cn/fuck/errors"
	logs "github.com/MicroOps-cn/fuck/log"
	w "github.com/MicroOps-cn/fuck/wrapper"
)

var (
	taskQueueMax = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Name: "job_task_queue_max",
		Help: "The maximum allowed number of task in the queue",
	}, nil)

	taskExecutedCount = prometheus.NewCounterFrom(
		stdprometheus.CounterOpts{
			Name: "job_task_executed_count",
			Help: "The number of completed tasks that have been executed",
		},
		[]string{"handler"},
	)
)

type Logger struct {
	kitlog.Logger
	Close func() error
}

type JobService interface {
	GetJobExecutorLogs(logId int32, logDateTime int64, from int) (content string, toLineNum int, isEnd bool, err error)
	GetJobExecutorLogger(logId int32, logDateTime int64) (logger *Logger, err error)
	RunJob(ctx context.Context, name string, jobId, logId int32, traceId string, handler func(ctx context.Context) error) error
	KillJob(ctx context.Context, jobId int32) error
	GetJobIdle(ctx context.Context, jobId int32) error
	Callback(ctx context.Context, logId int32, code int, msg string) error

	Register(ctx context.Context, name string, port int, root string) error
	Unregister(ctx context.Context) error

	GetJobHandler(ctx context.Context, name string) (handler func(ctx context.Context, params string) error)
	RegisterJobHandler(ctx context.Context, name string, handler func(ctx context.Context, params string) error) error

	CountRunningJob(ctx context.Context) int
	GetConcurrency() int

	runTask(ctx context.Context)
}

type JobEntry struct {
	Handler func(ctx context.Context, params string) error
	Name    string
}

var jobs []*JobEntry

func RegisterCronJobs(entities ...JobEntry) error {
	for _, entity := range entities {
		for _, job := range jobs {
			if job.Name == entity.Name {
				return errors2.New("duplicate job name")
			}
		}
		jobs = append(jobs, &entity)
	}
	return nil
}

func NewJobCron(ctx context.Context, cronConfigs []*config.JobCron, svc JobService) (*cron.Cron, error) {
	cr := cron.New()
	for idx, cronConfig := range cronConfigs {
		handler := svc.GetJobHandler(ctx, cronConfig.Name)
		if handler == nil {
			return nil, fmt.Errorf("the job `%s` not exists.", cronConfig.Name)
		}

		logger := logs.GetContextLogger(ctx)
		level.Info(logger).Log("msg", "add job to cron", "name", cronConfig.Name, "expr", cronConfig.Expr)
		if err := cr.AddFunc(cronConfig.Expr, func() {
			traceId := logs.NewTraceId()
			taskCtx, taskLogger := logs.NewContextLogger(ctx, logs.WithTraceId(traceId))
			err := svc.RunJob(taskCtx, cronConfig.Name, int32(idx), int32(rand.Int()), traceId, func(ctx context.Context) error {
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
		}); err != nil {
			cr.Stop()
			return nil, err
		}
	}
	return cr, nil
}

func NewJobService(ctx context.Context) (JobService, error) {
	jobConfig := config.Get().Job
	queueSize := 100
	taskQueueMax.Set(float64(queueSize))
	var jobSvc JobService
	taskQueue := w.NewRingQueue[Task](queueSize)
	switch backend := jobConfig.Scheduler.GetSchedulerBackend().(type) {
	case *config.JobOptions_Scheduler_XXLJob:
		client := xxljob.New(*backend.XXLJob)
		jobSvc = &jobService{c: client, taskQueue: taskQueue, register: client.Register, unregister: client.Unregister, jobs: slices.Clone(jobs)}
		level.Info(logs.GetContextLogger(ctx)).Log("msg", "run in xxljob runner", "concurrency", jobSvc.GetConcurrency())
	case *config.JobOptions_Scheduler_Local:
		localSchedule := &localJobSchedule{concurrency: runtime.NumCPU(), cron: cron.New()}
		jobSvc = &jobService{c: localSchedule, taskQueue: taskQueue, jobs: slices.Clone(jobs)}
		level.Info(logs.GetContextLogger(ctx)).Log("msg", "run in local runner", "concurrency", jobSvc.GetConcurrency())
		localSchedule.cron.Start()
		jobCron, err := NewJobCron(ctx, jobConfig.Cron, jobSvc)
		if err != nil {
			return nil, err
		}
		localSchedule.cron = jobCron
		jobCron.Start()
	default:
		return nil, errors.New("Unknown job scheduler")
	}
	for i := 0; i < jobSvc.GetConcurrency(); i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					jobSvc.runTask(ctx)
				}
			}
		}()
	}
	stdprometheus.MustRegister(
		stdprometheus.NewGaugeFunc(stdprometheus.GaugeOpts{
			Name: "job_task_queue_size",
			Help: "The current number of task in the queue.",
		}, func() float64 {
			return float64(taskQueue.Len())
		}),
		stdprometheus.NewGaugeFunc(stdprometheus.GaugeOpts{
			Name: "job_task_active",
			Help: "The number of tasks being executed",
		}, func() float64 {
			return float64(jobSvc.CountRunningJob(ctx))
		}),
	)

	return jobSvc, nil
}

type Task struct {
	task       func(ctx context.Context) error
	cancelFunc func()
	jobId      int32
	logId      int32
	traceId    string
	name       string
}

type JobClient interface {
	GetConcurrency() int
	Callback(ctx context.Context, id int32, code int, msg string) error
}

type NopAddFunc[T interface{}] struct {
}

type LogFile interface {
	io.Closer
	io.Reader
	io.Writer
}

type LogFS interface {
	OpenFile(name string, flag int, perm os.FileMode) (LogFile, error)
}

type jobService struct {
	c           JobClient
	taskQueue   *w.RingQueue[Task]
	runningTask []Task
	mux         sync.Mutex
	logDir      LogFS
	register    func(ctx context.Context, name string, port int, root string) error
	unregister  func(ctx context.Context) error
	jobs        []*JobEntry
}

func (s *jobService) GetJobHandler(ctx context.Context, name string) (handler func(ctx context.Context, params string) error) {
	for _, job := range s.jobs {
		if job.Name == name {
			return job.Handler
		}
	}
	return nil

}

func (s *jobService) RegisterJobHandler(ctx context.Context, name string, handler func(ctx context.Context, params string) error) error {
	for _, job := range jobs {
		if job.Name == name {
			return errors.New("duplicate job name")
		}
	}
	s.jobs = append(s.jobs, &JobEntry{Name: name, Handler: handler})
	return nil
}

func (s *jobService) Register(ctx context.Context, name string, port int, root string) error {
	if s.register == nil {
		return nil
	}
	return s.register(ctx, name, port, root)
}

func (s *jobService) Unregister(ctx context.Context) error {
	if s.unregister == nil {
		return nil
	}
	return s.unregister(ctx)
}

func (s *jobService) GetConcurrency() int {
	return s.c.GetConcurrency()
}

func (s *jobService) CountRunningJob(ctx context.Context) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return len(s.runningTask)
}

func (s *jobService) KillJob(ctx context.Context, jobId int32) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.runningTask = w.Filter(s.runningTask, func(item Task) bool {
		if item.jobId == jobId {
			item.cancelFunc()
			return false
		}
		return true
	})
	s.taskQueue.Remove(func(task Task) bool {
		if task.jobId == jobId {
			taskCtx, taskLogger := logs.NewContextLogger(ctx, logs.WithTraceId(task.traceId))
			if e := s.Callback(taskCtx, task.logId, 500, "user cancel"); e != nil {
				level.Error(taskLogger).Log("msg", "failed to callback job server", "err", e)
			}
			return true
		}
		return false
	})
	return nil
}

func (s *jobService) GetJobIdle(ctx context.Context, jobId int32) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	for _, task := range s.runningTask {
		if task.jobId == jobId {
			return errors.NewError(200, "busy", "500")
		}
	}
	for _, task := range s.taskQueue.List() {
		if task.jobId == jobId {
			return errors.NewError(200, "job executor is busy", "500")
		}
	}
	return nil
}

func (s *jobService) GetJobExecutorLogs(logId int32, logDateTime int64, from int) (content string, toLineNum int, isEnd bool, err error) {
	isEnd = true
	logFileName := fmt.Sprintf("job_%.0f_%d", math.Round(float64(logDateTime)/10000), logId)
	f, err := s.logDir.OpenFile(logFileName, os.O_RDONLY, 0o600)
	if err != nil {
		if os.IsNotExist(err) {
			return "", from, true, errors.NewError(404, err.Error())
		}
		return "", from, true, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	lineNum := 1
	for {
		line, prefix, err := reader.ReadLine()
		if len(line) > 0 {
			if from <= lineNum {
				content += string(line) + "\n"
			}
		} else {
			if errors2.Is(err, io.EOF) {
				return content, lineNum, true, nil
			}
			return content, lineNum, isEnd, err
		}

		if lineNum >= from+999 {
			isEnd = false
			return content, lineNum, isEnd, nil
		}
		if !prefix {
			lineNum++
		}
	}
}

func (s *jobService) GetJobExecutorLogger(logId int32, logDateTime int64) (*Logger, error) {
	logFileName := fmt.Sprintf("job_%.0f_%d", math.Round(float64(logDateTime)/10000), logId)
	f, err := s.logDir.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o600)
	if err != nil {
		return nil, err
	}
	logger := logs.WithCaller(6)(logs.New(logs.WithWriter(kitlog.NewSyncWriter(f))))
	return &Logger{Logger: logger, Close: f.Close}, nil
}

func (s *jobService) Callback(ctx context.Context, logId int32, code int, msg string) error {
	return s.c.Callback(ctx, logId, code, msg)
}

func (s *jobService) RunJob(ctx context.Context, name string, jobId, logId int32, traceId string, handler func(ctx context.Context) error) (err error) {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return s.taskQueue.Put(Task{jobId: jobId, name: name, logId: logId, traceId: traceId, task: handler})
	}
}

func (s *jobService) startTask(ctx context.Context) (context.Context, Task, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	task, err := s.taskQueue.Get()
	if err != nil {
		return nil, task, err
	}
	taskContext, cancelFunc := context.WithCancel(ctx)
	task.cancelFunc = cancelFunc
	s.runningTask = append(s.runningTask, task)
	return taskContext, task, nil
}

func (s *jobService) doneTask(jobId, logId int32) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.runningTask = w.Filter(s.runningTask, func(item Task) bool {
		return item.logId != logId || item.jobId != jobId
	})
}

func (s *jobService) runTask(ctx context.Context) {
	logger := logs.GetContextLogger(ctx)
	taskContext, task, err := s.startTask(ctx)
	if errors.Is(err, w.QueueNull) {
		time.Sleep(time.Millisecond * 100)
	} else if err != nil {
		level.Error(logger).Log("msg", "failed to get task", "err", err)
	} else {
		code := 200
		var msg string
		defer func() {
			if r := recover(); r != nil {
				msg = fmt.Sprintf("recover from exception: %s", r)
				code = 500
			}
			taskCtx, taskLogger := logs.NewContextLogger(ctx, logs.WithTraceId(task.traceId))
			if e := s.Callback(taskCtx, task.logId, code, msg); e != nil {
				level.Error(taskLogger).Log("msg", "failed to callback job server", "err", e)
			}
			taskExecutedCount.With("handler", task.name).Add(1)
		}()
		defer s.doneTask(task.jobId, task.logId)
		if task.task != nil {
			if err = task.task(taskContext); err != nil {
				code = 500
				msg = err.Error()
			}
		}
	}
}
