package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/MicroOps-cn/fuck/buffer"
	"github.com/MicroOps-cn/fuck/log"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
)

type BaseJobResponse struct {
	BaseResponse `json:",omitempty"`
	Code         int32  `json:"code,omitempty"`
	Msg          string `json:"msg,omitempty"`
}

func (r *BaseJobResponse) WithError(err error) *BaseJobResponse {
	code := "500"
	statusCode := 500
	if e, ok := err.(errors.ServerError); ok {
		statusCode = e.StatusCode()
		code = e.Code()
	}
	r.WithErrorMsg(statusCode, err.Error(), code)
	return r
}

func (r *BaseJobResponse) WithErrorMsg(statusCode int, err string, code ...string) *BaseJobResponse {
	r.Code = int32(statusCode)
	r.ErrorCode = strconv.Itoa(statusCode)
	if len(code) > 0 {
		if c, e := strconv.Atoi(code[0]); e == nil {
			r.Code = int32(c)
		}
		r.ErrorCode = code[0]
	}
	r.Msg = err
	r.Error = errors.NewServerError(statusCode, err, code...)
	return r
}

var compressBlackMatch = regexp.MustCompile("(?m) *\n *")

type RunJobRequest struct {
	JobId                 int32  `json:"jobId,omitempty"`
	ExecutorHandler       string `json:"executorHandler,omitempty"`
	ExecutorParams        string `json:"executorParams,omitempty"`
	ExecutorBlockStrategy string `json:"executorBlockStrategy,omitempty"`
	ExecutorTimeout       int32  `json:"executorTimeout,omitempty"`
	LogId                 int32  `json:"logId,omitempty"`
	LogDateTime           int64  `json:"logDateTime,omitempty"`
	GlueType              string `json:"glueType,omitempty"`
	GlueSource            string `json:"glueSource,omitempty"`
	GlueUpdateTime        int64  `json:"glueUpdatetime,omitempty"`
	BroadcastIndex        int32  `json:"broadcastIndex,omitempty"`
	BroadcastTotal        int32  `json:"broadcastTotal,omitempty"`
}

func (r *RunJobRequest) GetExecutorHandler() string {
	return r.ExecutorHandler
}

func (r *RunJobRequest) GetExecutorParams() string {
	return r.ExecutorParams
}

func MakeExecutorRunJobEndpoint(svc service.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (response BaseResponser, err error) {
		req := request.(Requester).GetRequestData().(*RunJobRequest)
		resp := &BaseJobResponse{BaseResponse: BaseResponse{Success: true, TraceId: log.GetTraceId(ctx)}, Code: 200}
		oriLogger := log.GetContextLogger(ctx)
		jobLogger, err := svc.GetJobExecutorLogger(req.LogId, req.LogDateTime)
		if err != nil {
			return resp.WithError(err), nil
		}
		executorHandler := req.GetExecutorHandler()
		handler := svc.GetJobHandler(ctx, executorHandler)
		if handler == nil {
			return resp.WithError(errors.NewServerError(404, "the handler is not found: "+req.GetExecutorHandler())), nil
		}
		params := req.GetExecutorParams()
		if err = svc.RunJob(ctx, req.GetExecutorHandler(), req.JobId, req.LogId, resp.TraceId, func(ctx context.Context) error {
			defer jobLogger.Close()
			logger := log.NewTeeLogger(log.NewTraceLogger(
				log.WithLogger(jobLogger),
				log.WithTraceId(resp.TraceId),
			), log.WithCaller(5)(oriLogger))
			ctx, _ = log.NewContextLogger(ctx, log.WithLogger(logger), log.WithTraceId(resp.TraceId))
			ctx, span := otel.GetTracerProvider().Tracer(global.AppName).Start(ctx, fmt.Sprintf("async run %s", executorHandler))
			defer span.End()
			span.SetAttributes(attribute.Int("job.log_id", int(req.LogId)), attribute.Int("job.id", int(req.JobId)))

			level.Info(logger).Log("msg", "execute job", log.WrapKeyName("handler"), executorHandler, log.WrapKeyName("params"), compressBlackMatch.ReplaceAllString(params, ""))
			if err = handler(ctx, params); err != nil {
				level.Error(logger).Log("msg", "failed to execute job", log.WrapKeyName("handler"), executorHandler, log.WrapKeyName("params"), compressBlackMatch.ReplaceAllString(params, ""), "err", err)
				span.SetStatus(codes.Error, err.Error())
			} else {
				level.Info(logger).Log("msg", "execute job finished")
				span.SetStatus(codes.Ok, "")
			}
			return err
		}); err != nil {
			return resp.WithError(err), nil
		}
		return resp, nil
	}
}

type JobIdleBeatRequest struct {
	JobId int32 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"jobId,omitempty"`
}

func MakeExecutorIdleBeatEndpoint(s service.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (response BaseResponser, err error) {
		req := request.(Requester).GetRequestData().(*JobIdleBeatRequest)
		baseResp := BaseJobResponse{Code: 200, BaseResponse: BaseResponse{Success: true, TraceId: log.GetTraceId(ctx)}}
		err = s.GetJobIdle(ctx, req.JobId)
		if err != nil {
			return baseResp.WithError(err), nil
		}
		return &baseResp, nil
	}
}

type KillJobRequest struct {
	JobId int32 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"jobId,omitempty"`
}

func MakeExecutorKillEndpoint(s service.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (response BaseResponser, err error) {
		req := request.(Requester).GetRequestData().(*KillJobRequest)
		baseResp := BaseJobResponse{Code: 200, BaseResponse: BaseResponse{Success: true, TraceId: log.GetTraceId(ctx)}}
		err = s.KillJob(ctx, req.JobId)
		if err != nil {
			return baseResp.WithError(err), nil
		}
		return &baseResp, nil
	}
}

func MakeExecutorBeatEndpoint(_ service.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (response BaseResponser, err error) {
		return &BaseJobResponse{BaseResponse: BaseResponse{Success: true, TraceId: log.GetTraceId(ctx)}, Code: 200}, nil
	}
}

type JobLogContent struct {
	FromLineNum int32  `protobuf:"varint,1,opt,name=from_line_num,json=fromLineNum,proto3" json:"fromLineNum,omitempty"`
	ToLineNum   int32  `protobuf:"varint,2,opt,name=to_line_num,json=toLineNum,proto3" json:"toLineNum,omitempty"`
	LogContent  string `protobuf:"bytes,3,opt,name=log_content,json=logContent,proto3" json:"logContent,omitempty"`
	IsEnd       bool   `protobuf:"varint,4,opt,name=is_end,json=isEnd,proto3" json:"isEnd,omitempty"`
}

type JobLogRequest struct {
	LogDateTim  int64 `protobuf:"varint,1,opt,name=log_date_tim,json=logDateTim,proto3" json:"logDateTim,omitempty"`
	LogId       int32 `protobuf:"varint,2,opt,name=log_id,json=logId,proto3" json:"logId,omitempty"`
	FromLineNum int32 `protobuf:"varint,3,opt,name=from_line_num,json=fromLineNum,proto3" json:"fromLineNum,omitempty"`
}

type JobLogResponse struct {
	BaseJobResponse      `protobuf:"bytes,1,opt,name=base_job_response,json=baseJobResponse,proto3,embedded=base_job_response" json:",omitempty"`
	Content              *JobLogContent `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func MakeExecutorLogEndpoint(svc service.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (response BaseResponser, err error) {
		req := request.(Requester).GetRequestData().(*JobLogRequest)
		baseResp := BaseJobResponse{Code: 200, BaseResponse: BaseResponse{Success: true, TraceId: log.GetTraceId(ctx)}}
		content, to, isEnd, err := svc.GetJobExecutorLogs(req.LogId, req.LogDateTim, int(req.FromLineNum))
		if err != nil {
			logger := log.GetContextLogger(ctx)
			level.Error(logger).Log("err", err, "msg", "failed to get executor log")
			if len(content) > 0 {
				isEnd = true
				content += "\n\n==================================================\nError: " + err.Error()
			} else {
				return baseResp.WithError(err), nil
			}
		}
		return &JobLogResponse{
			BaseJobResponse: baseResp,
			Content: &JobLogContent{
				LogContent:  content,
				FromLineNum: req.FromLineNum,
				ToLineNum:   int32(to),
				IsEnd:       isEnd,
			},
		}, nil
	}
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	logger := log.GetContextLogger(ctx)
	level.Error(logger).Log("err", err, "msg", "failed to http request")
	traceId := log.GetTraceId(ctx)
	resp := BaseJobResponse{
		BaseResponse: BaseResponse{
			ErrorMessage: err.Error(),
			TraceId:      traceId,
			Success:      false,
		},
		Msg: err.Error(),
	}
	if serverErr, ok := err.(errors.ServerError); ok {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(serverErr.StatusCode())
		resp.ErrorCode = serverErr.Code()
		if code, err := strconv.Atoi(serverErr.Code()); err == nil {
			resp.Code = int32(code)
		} else {
			resp.Code = 500
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		level.Info(logger).Log("msg", "failed to write response")
	}
}

func DecodeJobHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	logger := log.GetContextLogger(ctx)
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	logWriter := log.NewWriterAdapter(level.Debug(kitlog.With(logger, "resp", fmt.Sprintf("%#v", response), "caller", log.Caller(7))), log.Prefix("encoded http response: ", true))
	return json.NewEncoder(io.MultiWriter(w, buffer.LimitWriter(logWriter, 4096, buffer.LimitWriterIgnoreError))).Encode(response)
}
