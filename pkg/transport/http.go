package transport

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	errors2 "errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	stdlog "log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	fuckWebLogs "github.com/MicroOps-cn/fuck-web/pkg/logs"
	"github.com/MicroOps-cn/fuck/buffer"
	logs "github.com/MicroOps-cn/fuck/log"
	w "github.com/MicroOps-cn/fuck/wrapper"
	"github.com/asaskevich/govalidator"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	kitendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-openapi/spec"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/endpoint"
	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/httputil"
)

//go:embed static
var staticFs embed.FS

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.

func NewHTTPHandler(ctx context.Context, logger log.Logger, endpoints endpoint.Set, openapiPath string) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	m := restful.NewContainer()
	m.Filter(HTTPContextFilter(ctx))
	m.Filter(HTTPLoggingFilter(ctx))
	restful.TraceLogger(stdlog.New(log.NewStdlibAdapter(level.Info(logger)), "[restful]", stdlog.LstdFlags|stdlog.Lshortfile))
	var specTags []spec.Tag
	for _, serviceGenerator := range apiServiceSet {
		specTag, svcs := serviceGenerator(ctx, options, endpoints)
		for _, svc := range svcs {
			m.Add(svc)
		}
		specTags = append(specTags, specTag)
	}
	for _, serviceGenerator := range serviceGeneratorSet {
		specTag, svcs := serviceGenerator.WebServices(ctx, options, endpoints)
		for _, svc := range svcs {
			m.Add(svc)
		}
		specTags = append(specTags, specTag)
	}
	if openapiPath != "" {
		level.Info(logger).Log("msg", fmt.Sprintf("enable openapi on `%s`", openapiPath))
		specConf := restfulspec.Config{
			WebServices: m.RegisteredWebServices(),
			APIPath:     openapiPath,
			PostBuildSwaggerObjectHandler: func(swo *spec.Swagger) {
				swo.Info = &spec.Info{
					InfoProps: spec.InfoProps{
						Title:       "ItemTestService",
						Description: "Resource for managing ItemTests",
						Version:     "1.0.0",
					},
				}
				swo.Tags = specTags
			},
		}
		m.Add(restfulspec.NewOpenAPIService(specConf))
	}
	adminPrefix := ctx.Value(global.HTTPWebPrefixKey).(string)
	m.Handle(adminPrefix, http.StripPrefix(adminPrefix, NewStaticFileServer(ctx, w.M[fs.FS](fs.Sub(staticFs, "static")))))
	m.Handle("/healthz", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok\n"))
	}))
	return m
}

func SetCacheHeader(h http.Handler) http.Handler {
	modifyTime := time.Now().UTC().Truncate(time.Second)
	if executable, err := os.Executable(); err == nil {
		if stat, err := os.Stat(executable); err == nil {
			modifyTime = stat.ModTime().UTC()
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if clientModify := r.Header.Get("If-Modified-Since"); len(clientModify) > 0 {
			if clientModifyTime, err := time.Parse(http.TimeFormat, clientModify); err == nil {
				if !clientModifyTime.Before(modifyTime) {
					w.WriteHeader(304)
					return
				}
			}
		}
		w.Header().Set("Last-Modified", modifyTime.Format(http.TimeFormat))
		w.Header().Set("Cache-Control", "max-age=3600")
		h.ServeHTTP(w, r)
	})
}

type File struct {
	io.Reader
	stat StaticFSInfo
}

func (f File) Stat() (fs.FileInfo, error) {
	return f.stat, nil
}

func (f File) Read(p []byte) (int, error) {
	n, err := f.Reader.Read(p)
	return n, err
}

func (f File) Close() error {
	return nil
}

type StaticFSInfo struct {
	fs.FileInfo
	size int64
}

func (i StaticFSInfo) Size() int64 { return i.size }

type StaticFS struct {
	fs            fs.FS
	indexTmpl     *template.Template
	indexTmplStat fs.FileInfo
	ctx           context.Context
}

func (t *StaticFS) GetIndexFile() (fs.File, error) {
	gConf := config.Get().GetGlobal()
	var buf bytes.Buffer
	err := t.indexTmpl.Execute(&buf, map[string]interface{}{
		"title":     gConf.Title,
		"sub_title": gConf.SubTitle,
		"logo":      gConf.Logo,
	})
	if err != nil {
		return nil, err
	}
	return &File{Reader: &buf, stat: StaticFSInfo{FileInfo: t.indexTmplStat, size: int64(buf.Len())}}, nil
}

func (t *StaticFS) Open(name string) (fs.File, error) {
	if name == "index.html" {
		return t.GetIndexFile()
	}
	open, err := t.fs.Open(name)
	if errors2.Is(err, fs.ErrNotExist) {
		return t.GetIndexFile()
	}
	return open, err
}

func NewStaticFileServer(ctx context.Context, fileSystem fs.FS) http.Handler {
	f, err := fileSystem.Open("index.html")
	if err != nil {
		panic(fmt.Errorf("failed to open index file: %s", err))
	}
	stat, err := f.Stat()
	if err != nil {
		panic(fmt.Errorf("failed to get index file stat: %s", err))
	}
	tmpl, err := template.ParseFS(fileSystem, "index.html")
	if err != nil {
		panic(fmt.Errorf("failed to open index file: %s", err))
	}
	return SetCacheHeader(http.FileServer(http.FS(&StaticFS{ctx: ctx, fs: fileSystem, indexTmpl: tmpl, indexTmplStat: stat})))
}

const DisableStackTrace = "__disable_stack_trace__"

func errorEncoder(ctx context.Context, err error, writer http.ResponseWriter) {
	logger := logs.GetContextLogger(ctx)

	if ctx.Value(DisableStackTrace) != true && logs.DefaultLoggerConfig.Level.String() == "debug" {
		level.Error(logs.WithPrint(w.NewStringer(func() string {
			return fmt.Sprintf("%+v", err)
		}))(logger)).Log("msg", "failed to http request", "err", err)
	} else {
		level.Error(logger).Log("msg", "failed to http request", "err", err)
	}

	traceId := logs.GetTraceId(ctx)
	resp := responseWrapper{
		ErrorMessage: err.Error(),
		TraceId:      traceId,
		Success:      false,
	}
	if serverErr, ok := err.(errors.ServerError); ok {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(serverErr.StatusCode())
		resp.ErrorCode = serverErr.Code()
	} else {
		resp.ErrorCode = strconv.Itoa(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	if err = json.NewEncoder(writer).Encode(resp); err != nil {
		level.Info(logger).Log("msg", "failed to write response", "err", err)
	}
}

type ResponseWrapper[T any] struct {
	Data T `json:"data"`
}

type responseWrapper struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorCode    string      `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	TraceId      string      `json:"traceId"`
	Current      int64       `json:"current,omitempty"`
	PageSize     int64       `json:"pageSize,omitempty"`
	Total        int64       `json:"total"`
}

type HTTPRequest[T any] struct {
	Data            T `json:"data"`
	restfulRequest  *restful.Request
	restfulResponse *restful.Response
}

func (b HTTPRequest[T]) GetRequestData() interface{} {
	return &b.Data
}

func (b HTTPRequest[T]) GetRestfulRequest() *restful.Request {
	return b.restfulRequest
}

func (b HTTPRequest[T]) GetRestfulResponse() *restful.Response {
	return b.restfulResponse
}

var _ endpoint.RestfulRequester = &HTTPRequest[any]{}

func isProtoMessage(v interface{}) (proto.Message, bool) {
	msg, ok := v.(proto.Message)
	return msg, ok
}

func valid(data interface{}) (bool, error) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Struct:
		return govalidator.ValidateStruct(data)
	case reflect.Slice:
		valOf := reflect.ValueOf(data)
		for i := 0; i < valOf.Len(); i++ {
			b, err := valid(valOf.Index(i).Interface())
			if err != nil || !b {
				return b, err
			}
		}
	}
	return true, nil
}

// DecodeHTTPRequest Decode HTTP requests into request types
func DecodeHTTPRequest[RequestType any](_ context.Context, stdReq *http.Request) (interface{}, error) {
	restfulReq := stdReq.Context().Value(global.RestfulRequestContextName).(*restful.Request)
	restfulResp := stdReq.Context().Value(global.RestfulResponseContextName).(*restful.Response)
	hasSensitiveData, _ := stdReq.Context().Value(global.MetaSensitiveData).(bool)
	req := HTTPRequest[RequestType]{restfulRequest: restfulReq, restfulResponse: restfulResp}
	var err error
	logger := logs.GetContextLogger(stdReq.Context())
	r := restfulReq.Request
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" || r.Method == "DELETE" {
		contentType := httputil.GetContentType(r.Header)
		if contentType == "multipart/form-data" {
			restfulReq.Request.Body = http.MaxBytesReader(restfulResp.ResponseWriter, r.Body, int64(config.Get().Global.MaxUploadSize))
			if err = restfulReq.Request.ParseMultipartForm(int64(config.Get().Global.MaxBodySize)); err != nil {
				return nil, errors.NewServerError(http.StatusBadRequest, "request too large")
			}
		} else {
			r.Body = http.MaxBytesReader(restfulResp.ResponseWriter, r.Body, int64(config.Get().Global.MaxBodySize))
			if contentType == "application/x-www-form-urlencoded" {
				if err = r.ParseForm(); err != nil {
					return nil, errors.WithServerError(http.StatusBadRequest, err, "failed to parse form data")
				} else if len(r.Form) > 0 {
					if err = httputil.UnmarshalURLValues(r.Form, &req.Data); err != nil {
						return nil, errors.WithServerError(http.StatusBadRequest, err, fmt.Sprintf("failed to decode form data: data=%s", r.Form))
					}
				}
			} else if contentType == restful.MIME_JSON {
				if data, ok := isProtoMessage(&req.Data); ok {
					var reader io.Reader = r.Body
					if !hasSensitiveData {
						logWriter := logs.NewWriterAdapter(level.Debug(logs.WithCaller(12)(logger)), logs.Prefix("decode http request: ", true))
						reader = io.TeeReader(r.Body, buffer.LimitWriter(logWriter, 1024, buffer.LimitWriterIgnoreError))
					}
					if err = jsonpb.Unmarshal(reader, data); err != nil {
						return nil, errors.WithServerError(http.StatusBadRequest, err, "failed to decode request body")
					}
				} else {
					var reader io.Reader = r.Body
					if !hasSensitiveData {
						logWriter := logs.NewWriterAdapter(level.Debug(logs.WithCaller(9)(logger)), logs.Prefix("decode http request: ", true))
						reader = io.TeeReader(r.Body, buffer.LimitWriter(logWriter, 1024, buffer.LimitWriterIgnoreError))
					}
					if err = json.NewDecoder(reader).Decode(&req.Data); err != nil {
						return nil, errors.WithServerError(http.StatusBadRequest, err, "failed to decode request body")
					}
				}
			}
		}
	}

	query := restfulReq.Request.URL.Query()
	if len(query) > 0 {
		if err = httputil.UnmarshalURLValues(query, &req.Data); err != nil {
			return nil, errors.WithServerError(http.StatusBadRequest, err, "failed to decode url query")
		}
	}
	if len(restfulReq.PathParameters()) > 0 {
		if err = httputil.UnmarshalURLValues(httputil.MapToURLValues(restfulReq.PathParameters()), &req); err != nil {
			return nil, errors.WithServerError(http.StatusBadRequest, err, "failed to decode path parameters")
		}
	}

	req.restfulRequest = restfulReq
	req.restfulResponse = restfulResp
	level.Debug(logger).Log("msg", "decoded http request", fuckWebLogs.WrapKeyName("data"), w.JSONStringer(req.Data), fuckWebLogs.WrapKeyName("type"), fmt.Sprintf("%T", req.Data))
	if ok, err := valid(req.Data); err != nil {
		return &req, errors.NewServerError(http.StatusBadRequest, err.Error())
	} else if !ok {
		return &req, errors.NewServerError(http.StatusBadRequest, "params error")
	}

	return &req, err
}

// encodeHTTPResponse Encode the response as an HTTP response message
func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response == httputil.NopResponse {
		return nil
	}
	logger := logs.GetContextLogger(ctx)
	valOf := reflect.ValueOf(response)
	if valOf.Kind() != reflect.Ptr {
		errorEncoder(ctx, errors.NewServerError(500, fmt.Sprintf("result data must be a pointer")), w)
		return nil
	}
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	traceId := logs.GetTraceId(ctx)
	resp := responseWrapper{Success: true, TraceId: traceId}
	if l, ok := response.(endpoint.Lister); ok {
		resp.Data = l.GetData()
		resp.Total = l.GetTotal()
		resp.PageSize = l.GetPageSize()
		resp.Current = l.GetCurrent()
	} else if response != nil {
		if t, ok := response.(endpoint.Total); ok {
			resp.Total = t.GetTotal()
		}
		if t, ok := response.(endpoint.HasData); ok {
			resp.Data = t.GetData()
		} else if _, ok := response.(endpoint.BaseResponser); !ok {
			resp.Data = response
		}
	}

	if ctx.Value(global.MetaSensitiveData) == true {
		return json.NewEncoder(w).Encode(resp)
	}
	logWriter := logs.NewWriterAdapter(level.Debug(log.With(logs.WithCaller(7)(logger), "resp", fmt.Sprintf("%#v", resp))), logs.Prefix("encoded http response: ", true))
	return json.NewEncoder(io.MultiWriter(w, buffer.LimitWriter(logWriter, 1024, buffer.LimitWriterIgnoreError))).Encode(resp)
}

func SimpleEncodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	logger := logs.GetContextLogger(ctx)
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	logWriter := logs.NewWriterAdapter(level.Debug(log.With(logs.WithCaller(7)(logger), "resp", fmt.Sprintf("%#v", response))), logs.Prefix("encoded http response: ", true))
	return json.NewEncoder(io.MultiWriter(w, buffer.LimitWriter(logWriter, 1024, buffer.LimitWriterIgnoreError))).Encode(response)
}

func NoopEncodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
	}
	return nil
}

func WrapHTTPHandler(pctx context.Context, h *httptransport.Server) func(*restful.Request, *restful.Response) {
	return func(req *restful.Request, resp *restful.Response) {
		ctx := req.Request.Context()
		if ctx == nil {
			ctx = pctx
		}
		request := req.Request.WithContext(context.WithValue(context.WithValue(ctx, global.RestfulResponseContextName, resp), global.RestfulRequestContextName, req))
		h.ServeHTTP(resp, request)
	}
}

func NewKitHTTPServer[RequestType any](ctx context.Context, dp kitendpoint.Endpoint, options []httptransport.ServerOption) restful.RouteFunction {
	return WrapHTTPHandler(ctx, httptransport.NewServer(
		dp,
		DecodeHTTPRequest[RequestType],
		encodeHTTPResponse,
		options...,
	))
}

func NewSimpleKitHTTPServer[RequestType any](
	ctx context.Context,
	dp kitendpoint.Endpoint,
	dec httptransport.DecodeRequestFunc,
	enc httptransport.EncodeResponseFunc, options []httptransport.ServerOption,
) restful.RouteFunction {
	return WrapHTTPHandler(ctx, httptransport.NewServer(
		dp,
		dec,
		enc,
		options...,
	))
}

const QueryTypeKey = "__query_type__"

func NewWebService(rootPath string, gv schema.GroupVersion, doc string) *restful.WebService {
	webservice := restful.WebService{}
	webservice.Path(rootPath + "/" + gv.Version + "/" + gv.Group).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).Doc(doc)
	return &webservice
}

func NewSimpleWebService(rootPath string, doc string) *restful.WebService {
	webservice := restful.WebService{}
	webservice.Path(rootPath).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).Doc(doc)
	return &webservice
}

const RootPath = "/api"

func StructToQueryParams(obj interface{}, nameFilter ...string) []*restful.Parameter {
	var params []*restful.Parameter
	typeOfObj := reflect.TypeOf(obj)
	valueOfObj := reflect.ValueOf(obj)
	// 通过 #NumField 获取结构体字段的数量
loopObjFields:
	for i := 0; i < typeOfObj.NumField(); i++ {
		field := typeOfObj.Field(i)

		if field.Type.Kind() == reflect.Struct && field.Anonymous {
			params = append(params, StructToQueryParams(valueOfObj.Field(i).Interface(), nameFilter...)...)
		} else {
			if len(nameFilter) > 0 {
				for _, name := range nameFilter {
					if name == field.Name {
						goto handleField
					}
				}
				continue loopObjFields
			}
		handleField:
			jsonTag := strings.Split(field.Tag.Get("json"), ",")
			if len(jsonTag) > 0 && jsonTag[0] != "-" && jsonTag[0] != "" {
				param := restful.QueryParameter(
					jsonTag[0],
					field.Tag.Get("description"),
				).DataType(field.Type.String())
				if len(jsonTag) > 1 && jsonTag[1] == "omitempty" {
					param.Required(false)
				} else {
					param.Required(true)
				}
				if tag := field.Tag.Get("enum"); tag != "" {
					enums := map[string]string{}
					for idx, s := range strings.Split(tag, "|") {
						enums[strconv.Itoa(idx)] = s
					}
					param.AllowableValues(enums)
				} else if protoTag := field.Tag.Get("protobuf"); protoTag != "" {
					var typeName string
					for _, s := range strings.Split(protoTag, ",") {
						if strings.HasPrefix(s, "enum=") {
							typeName = s[5:]
							break
						}
					}
					if len(typeName) != 0 {
						enumMap := proto.EnumValueMap(typeName)
						enums := make(map[string]string, len(enumMap))
						for v, idx := range enumMap {
							enums[strconv.Itoa(int(idx))] = v
						}
						param.AllowableValues(enums)
						param.AddExtension("$ref", typeName)
						param.DataType("string")
						param.DataFormat("string")
					}
				}
				params = append(params, param)
			}
		}
	}
	return params
}
