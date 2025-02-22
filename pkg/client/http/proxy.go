package http

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/MicroOps-cn/fuck/log"
	w "github.com/MicroOps-cn/fuck/wrapper"
	log2 "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var DefaultProxyClient = NewProxyClient()

func NewProxyClient() *http.Client {
	mux := sync.Mutex{}
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &dynamicServerHostTransport{
			mux: &mux,
			tr: otelhttp.NewTransport(&http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					isInsecureSkipVerify, _ := ctx.Value(optionInsecureSkipVerify{}).(bool)
					transparentServerName, _ := ctx.Value(optionTransparentServerName{}).(string)
					caCert, _ := ctx.Value(optionCaCert{}).(*x509.CertPool)
					conn, err := tls.Dial(network, addr, &tls.Config{
						ServerName:         transparentServerName,
						InsecureSkipVerify: isInsecureSkipVerify,
						RootCAs:            caCert,
					})
					if err != nil {
						return nil, err
					}
					return conn, conn.HandshakeContext(ctx)
				},
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			}),
		},
	}
}

type dynamicServerHostTransport struct {
	tr  http.RoundTripper
	mux *sync.Mutex
}

type optionInsecureSkipVerify struct{}

func WithInsecureSkipVerify(req *http.Request) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), optionInsecureSkipVerify{}, true))
}

type optionCaCert struct{}

func WithCaCert(req *http.Request, caCert *x509.CertPool) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), optionCaCert{}, caCert))
}

type optionTransparentServerName struct{}

func WithTransparentServerName(req *http.Request, serverName string) *http.Request {
	req.Host = serverName
	return req.WithContext(context.WithValue(req.Context(), optionTransparentServerName{}, serverName))
}

func (d dynamicServerHostTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return d.tr.RoundTrip(req)
}

func SendProxyRequest(req *http.Request) (resp *http.Response, err error) {
	start := time.Now()
	defer func() {
		isInsecureSkipVerify, _ := req.Context().Value(optionInsecureSkipVerify{}).(bool)

		logger := log2.With(log.GetContextLogger(req.Context()),
			"msg", "sed proxy request",
			"host", req.Host,
			"target", req.URL,
			"insecureSkipVerify", isInsecureSkipVerify,
			"duration", time.Since(start),
		)
		if caCert, ok := req.Context().Value(optionCaCert{}).(*x509.CertPool); ok {
			logger = log2.With(logger, "caCert", strings.Join(w.Map(caCert.Subjects(), func(item []byte) string {
				return string(item)
			}), ","))
		}
		if err != nil {
			level.Debug(logger).Log("err", err)
		} else {
			level.Debug(logger).Log("code", resp.StatusCode)
		}
	}()
	return DefaultProxyClient.Do(req)
}
