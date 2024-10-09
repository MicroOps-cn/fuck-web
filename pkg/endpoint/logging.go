package endpoint

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
)

func MakeGetEventsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Requester).GetRequestData().(*GetEventsRequest)
		resp := NewBaseListResponse[[]*models.Event](&req.BaseListRequest)
		startTime, err := time.Parse(time.RFC3339Nano, req.StartTime)
		if err != nil {
			return nil, errors.NewServerError(http.StatusBadRequest, fmt.Sprintf("Parameter Error: failed to parse startTime: %s", err))
		}
		endTime, err := time.Parse(time.RFC3339Nano, req.EndTime)
		if err != nil {
			return nil, errors.NewServerError(http.StatusBadRequest, fmt.Sprintf("Parameter Error: failed to parse endTime: %s", err))
		}
		resp.Total, resp.Data, resp.Error = svc.GetEvents(ctx, map[string]string{}, req.Keywords, startTime, endTime, req.Current, req.PageSize)
		return &resp, nil
	}
}

func MakeGetEventLogsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Requester).GetRequestData().(*GetEventLogsRequest)
		resp := NewBaseListResponse[[]*models.EventLog](&req.BaseListRequest)
		resp.Total, resp.Data, resp.Error = svc.GetEventLogs(ctx, map[string]string{"event_id": req.EventId}, req.Keywords, req.Current, req.PageSize)
		return &resp, nil
	}
}

func MakeGetCurrentUserEventsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Requester).GetRequestData().(*GetCurrentUserEventsRequest)
		resp := NewBaseListResponse[[]*models.Event](&req.BaseListRequest)
		user, ok := ctx.Value(global.MetaUser).(*models.User)
		if !ok || user == nil {
			resp.Error = errors.NotLoginError()
			return &resp, nil
		}
		startTime, err := time.Parse(time.RFC3339Nano, req.StartTime)
		if err != nil {
			return nil, errors.NewServerError(http.StatusBadRequest, fmt.Sprintf("Parameter Error: failed to parse startTime: %s", err))
		}
		endTime, err := time.Parse(time.RFC3339Nano, req.EndTime)
		if err != nil {
			return nil, errors.NewServerError(http.StatusBadRequest, fmt.Sprintf("Parameter Error: failed to parse endTime: %s", err))
		}
		resp.Total, resp.Data, resp.Error = svc.GetEvents(ctx, map[string]string{"user_id": user.Id}, req.Keywords, startTime, endTime, req.Current, req.PageSize)
		return &resp, nil
	}
}

func MakeGetCurrentUserEventLogsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Requester).GetRequestData().(*GetCurrentUserEventLogsRequest)
		resp := NewBaseListResponse[[]*models.EventLog](&req.BaseListRequest)
		user, ok := ctx.Value(global.MetaUser).(*models.User)
		if !ok || user == nil {
			resp.Error = errors.NotLoginError()
			return &resp, nil
		}
		resp.Total, resp.Data, resp.Error = svc.GetEventLogs(ctx, map[string]string{"event_id": req.EventId, "user_id": user.Id}, req.Keywords, req.Current, req.PageSize)
		return &resp, nil
	}
}
