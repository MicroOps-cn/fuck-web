package endpoint

import (
	"context"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
)

type BaseResponser interface {
	GetTraceId() string
	GetSuccess() bool
	GetErrorCode() string
	GetErrorMessage() string
	Failed() error
	GetData() interface{}
}

type Lister interface {
	GetPageSize() int64
	GetCurrent() int64
	GetTotal() int64
	GetData() interface{}
}
type HasData interface {
	GetData() interface{}
}
type Total interface {
	GetTotal() int64
}

type Requester interface {
	GetRequestData() interface{}
}

type RestfulRequester interface {
	Requester
	GetRestfulRequest() *restful.Request
	GetRestfulResponse() *restful.Response
}

func (l BaseResponse) Failed() error {
	if len(l.ErrorMessage) != 0 {
		return errors.WithServerError(http.StatusOK, l.Error, l.ErrorMessage, l.ErrorCode)
	}
	return l.Error
}

func NewBaseListResponse[T any](req *BaseListRequest) ListResponseWrapper[T] {
	if req == nil {
		req = &BaseListRequest{}
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	if req.Current == 0 {
		req.Current = 1
	}
	return ListResponseWrapper[T]{
		BaseListResponse: BaseListResponse{
			PageSize: req.PageSize,
			Current:  req.Current,
			BaseTotalResponse: BaseTotalResponse{
				BaseResponse: BaseResponse{},
			},
		},
	}
}

type SimpleResponseWrapper[DataType any] struct {
	BaseResponse
	Data DataType `json:"data"`
}

func (l SimpleResponseWrapper[T]) GetData() interface{} {
	return l.Data
}

type TotalResponseWrapper[DataType any] struct {
	BaseTotalResponse
	Data DataType `json:"data"`
}

func (l TotalResponseWrapper[T]) GetData() interface{} {
	return l.Data
}

type ListResponseWrapper[DataType any] struct {
	BaseListResponse
	Data DataType `json:"data"`
}

func (l ListResponseWrapper[T]) GetData() interface{} {
	return l.Data
}

func (l BaseResponse) GetData() interface{} {
	return nil
}

type Endpoint func(ctx context.Context, request interface{}) (response BaseResponser, err error)
