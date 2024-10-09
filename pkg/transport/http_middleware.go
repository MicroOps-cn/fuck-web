package transport

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck/buffer"
	http2 "github.com/MicroOps-cn/fuck/http"
	"github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/signals"
	w "github.com/MicroOps-cn/fuck/wrapper"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/go-kit/kit/metrics/prometheus"
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/MicroOps-cn/fuck-web/pkg/endpoint"
	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
	"github.com/MicroOps-cn/fuck-web/pkg/logs"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/httputil"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/sign"
)

func getTokenByRequest(req *http.Request) *endpoint.GetSessionParams {
	loginSessionID, err := req.Cookie(global.LoginSession)
	if err == nil {
		return &endpoint.GetSessionParams{
			Token:     loginSessionID.Value,
			TokenType: models.TokenTypeLoginSession,
		}
	}
	if auth := req.Header.Get("Authorization"); len(auth) != 0 {
		if strings.HasPrefix(auth, "Bearer ") {
			return &endpoint.GetSessionParams{
				Token:     strings.TrimPrefix(auth, "Bearer "),
				TokenType: models.TokenTypeToken,
			}
		}
	}
	return nil
}

func getAuthReqByRequest(req *http.Request) (*HTTPRequest[endpoint.AuthenticationRequest], error) {
	var err error
	authReq := &HTTPRequest[endpoint.AuthenticationRequest]{}
	if username, password, ok := req.BasicAuth(); ok {
		authReq.Data.AuthKey = username
		authReq.Data.AuthSecret = password
	} else {
		query := req.URL.Query()
		if query.Get("authKey") != "" {
			if err = httputil.UnmarshalURLValues(query, &authReq); err != nil {
				return nil, errors.NewServerError(400, "unknown exception")
			}
		}
	}
	if len(authReq.Data.AuthKey) > 0 || len(authReq.Data.AuthSecret) > 0 {
		if authReq.Data.AuthSign != "" {
			if authReq.Data.Payload, err = sign.GetPayloadFromHTTPRequest(req); err != nil {
				return nil, errors.ParameterError("Failed to get payload")
			}
		}
		return authReq, nil
	}
	return nil, nil
}

func HTTPAuthenticationFilter(endpoints endpoint.Set) restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, filterChan *restful.FilterChain) {
		ctx := req.Request.Context()
		var authError error
		if req.SelectedRoute() == nil {
			errorEncoder(ctx, errors.NewServerError(http.StatusNotFound, "Not Found: "+req.Request.RequestURI), resp)
			return
		}
		needLogin, ok := ctx.Value(global.MetaNeedLogin).(bool)
		if !ok {
			needLogin = true
		}
		if token := getTokenByRequest(req.Request); token != nil {
			ctx = context.WithValue(ctx, global.LoginSession, token.Token)
			req.Request = req.Request.WithContext(ctx)
			sessionReq := &HTTPRequest[endpoint.GetSessionParams]{restfulRequest: req, restfulResponse: resp, Data: *token}
			if user, err := endpoints.GetSessionByToken(ctx, sessionReq); err == nil {
				if u := user.(*models.User); u != nil {
					if passwordExpireTime := config.GetRuntimeConfig().Security.PasswordExpireTime; passwordExpireTime > 0 && needLogin {
						if u.ExtendedData != nil && time.Since(u.ExtendedData.PasswordModifyTime) > time.Duration(passwordExpireTime)*time.Hour*24 {
							_, _ = endpoints.UserLogout(ctx, HTTPRequest[any]{restfulRequest: req, restfulResponse: resp})
							errorEncoder(ctx, errors.NewServerError(http.StatusOK, "Your password has expired. Please change the password and log in again.", errors.CodeUserNeedResetPassword), resp)
							return
						}
					}
					ctx = context.WithValue(ctx, global.MetaUser, user)
					req.Request = req.Request.WithContext(ctx)
					filterChan.ProcessFilter(req, resp)
					return
				}
				authError = errors.NewServerError(http.StatusUnauthorized, "can't get user by token")
			} else {
				authError = errors.WithServerError(http.StatusUnauthorized, err, "failed to get session by token")
			}
		}
		if authReq, err := getAuthReqByRequest(req.Request); err != nil {
			errorEncoder(ctx, err, resp)
			return
		} else if authReq != nil {
			if user, err := endpoints.Authentication(ctx, authReq); err == nil {
				if u := user.(*models.User); u != nil {
					if passwordExpireTime := config.GetRuntimeConfig().Security.PasswordExpireTime; passwordExpireTime > 0 && needLogin {
						if u.ExtendedData != nil && time.Since(u.ExtendedData.PasswordModifyTime) > time.Duration(passwordExpireTime)*time.Hour*24 {
							errorEncoder(ctx, errors.NewServerError(http.StatusOK, "Your password has expired. Please change the password and log in again.", errors.CodeUserNeedResetPassword), resp)
							return
						}
					}
					req.Request = req.Request.WithContext(context.WithValue(ctx, global.MetaUser, user))
					filterChan.ProcessFilter(req, resp)
					return
				}
			} else if err != nil {
				authError = errors.WithServerError(http.StatusUnauthorized, err, "failed to get session by auth request")
			}
		}

		if !needLogin {
			filterChan.ProcessFilter(req, resp)
			return
		}
		if autoRedirectToLoginPage, ok := ctx.Value(global.MetaAutoRedirectToLoginPage).(bool); ok && autoRedirectToLoginPage {
			redirectURI := req.Request.RequestURI
			if externalURL, ok := ctx.Value(global.HTTPExternalURLKey).(string); ok {
				extURL, err := url.Parse(externalURL)
				if err == nil {
					extURL.Path = http2.JoinPath(extURL.Path, req.Request.URL.Path)
					extURL.RawQuery = req.Request.URL.RawQuery
					redirectURI = extURL.String()
				}
			}
			if loginURL, ok := ctx.Value(global.HTTPLoginURLKey).(string); ok && len(loginURL) > 0 {
				resp.Header().Set("Location", fmt.Sprintf("%s?redirect_uri=%s", loginURL, url.QueryEscape(redirectURI)))
			} else {
				resp.Header().Set("Location", fmt.Sprintf("/admin/account/login?redirect_uri=%s", url.QueryEscape(redirectURI)))
			}
			resp.WriteHeader(302)
			return
		}
		if authError != nil {
			errorEncoder(ctx, errors.WithServerError(http.StatusUnauthorized, authError, "Not logged in or identity expired"), resp)
		} else {
			errorEncoder(context.WithValue(ctx, DisableStackTrace, true), errors.NewServerError(http.StatusUnauthorized, "Not logged in or identity expired"), resp)
		}
	}
}

func getSafeHeader(req *http.Request) fmt.Stringer {
	header := req.Header.Clone()
	cookies := req.Cookies()
	return w.NewStringer(func() string {
		if auth := header.Get("Authorization"); len(auth) > 0 {
			header.Set("Authorization", fmt.Sprintf("[sha256]%x", sha256.Sum256([]byte(auth))))
		}
		header.Del("Cookie")
		for _, cookie := range cookies {
			cookieVal := cookie.Value
			if cookie.Name == global.LoginSession {
				cookieVal = fmt.Sprintf("[sha256]%x", sha256.Sum256([]byte(cookie.Value)))
			}
			header.Add("Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookieVal))
		}
		return w.JSONStringer(header).String()
	})
}

var (
	requestsTotal = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Tracks the number of HTTP requests.",
	}, []string{"method", "code", "api"})
	requestDuration = prometheus.NewHistogramFrom(
		stdprometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Tracks the latencies for HTTP requests.",
			Buckets: stdprometheus.ExponentialBuckets(0.1, 3, 5),
		},
		[]string{"method", "code", "api"},
	)
)

func HTTPContextFilter(pctx context.Context) restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		ch := signals.SignalHandler()
		ch.AddRequest(1)
		defer ch.DoneRequest()

		var route restful.RouteReader
		start := time.Now()
		defer func() {
			if route != nil {
				requestsTotal.With("method", route.Method(), "code", strconv.Itoa(resp.StatusCode()), "api", route.Path()).Add(1)
				requestDuration.With("method", route.Method(), "code", strconv.Itoa(resp.StatusCode()), "api", route.Path()).
					Observe(float64(time.Since(start) / time.Second))
			}
		}()
		ctx := req.Request.Context()
		if ctx == nil {
			ctx = pctx
		}
		if route = req.SelectedRoute(); route != nil && route.Metadata() != nil {
			metadata := req.SelectedRoute().Metadata()
			for key, val := range metadata {
				ctx = context.WithValue(ctx, key, val)
			}
		}
		req.Request = req.Request.WithContext(ctx)
		chain.ProcessFilter(req, resp)
	}
}

func HTTPLoggingFilter(pctx context.Context) func(req *restful.Request, resp *restful.Response, filterChan *restful.FilterChain) {
	return func(req *restful.Request, resp *restful.Response, filterChan *restful.FilterChain) {
		var logger kitlog.Logger
		ctx := req.Request.Context()
		if ctx == nil {
			ctx = pctx
		}
		hasSensitiveData, _ := ctx.Value(global.MetaSensitiveData).(bool)
		start := time.Now()
		spanName := req.Request.RequestURI
		var spanOptions []trace.SpanStartOption
		if req.SelectedRoute() != nil {
			spanName = req.SelectedRoute().Operation()
		}
		ctx, span := otel.GetTracerProvider().Tracer(config.Get().GetAppName()).Start(ctx, spanName, spanOptions...)
		traceId := span.SpanContext().TraceID()
		traceIdStr := traceId.String()
		if !traceId.IsValid() {
			traceIdStr = log.NewTraceId()
		}

		ctx, _ = log.NewContextLogger(ctx, log.WithTraceId(traceIdStr))
		req.Request = req.Request.WithContext(ctx)
		logger = log.GetContextLogger(ctx)
		defer func() {
			if r := recover(); r != nil {
				span.SetStatus(codes.Error, fmt.Sprintf("%+v", r))
				errorEncoder(ctx, errors.NewServerError(http.StatusInternalServerError, "Server exception"), resp)
				buf := bytes.NewBufferString(fmt.Sprintf("recover from panic situation: - %v\n", r))
				for i := 2; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					buf.WriteString(fmt.Sprintf("    %s:%d\n", file, line))
				}
				level.Error(logger).Log("msg", buf.String())
			}
			logger = kitlog.With(logger,
				"msg", "HTTP response send.",
				logs.TitleKey, "response",
				logs.WrapKeyName("httpURI"), req.Request.RequestURI,
				logs.WrapKeyName("status"), resp.StatusCode(),
				logs.WrapKeyName("contentType"), resp.Header().Get("Content-Type"),
				logs.WrapKeyName("contentLength"), resp.ContentLength(),
			)
			level.Info(logger).Log(logs.WrapKeyName("totalTime"), fmt.Sprintf("%dms", time.Since(start).Milliseconds()))
			span.End()
		}()
		var reqBody fmt.Stringer
		if !hasSensitiveData {
			preBuf := w.M[buffer.PreReader](buffer.NewPreReader(req.Request.Body, 1024))
			req.Request.Body = preBuf
			reqBody = preBuf
		} else {
			reqBody = bytes.NewBufferString("<body>")
		}
		level.Info(logger).Log(
			"msg", "HTTP request received.",
			logs.TitleKey, "request",
			logs.WrapKeyName("httpURI"), fmt.Sprintf("%s %s %s", req.Request.Method, req.Request.RequestURI, req.Request.Proto),
			logs.WrapKeyName("contentType"), req.HeaderParameter("Content-Type"),
			logs.WrapKeyName("header"), getSafeHeader(req.Request),
			logs.WrapKeyName("contentLength"), req.Request.ContentLength,
			logs.WrapKeyName("body"), reqBody,
		)
		filterChan.ProcessFilter(req, resp)
	}
}
