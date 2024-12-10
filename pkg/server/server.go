package server

import (
	"context"
	"fmt"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	//revive:disable:blank-imports
	_ "net/http/pprof"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/endpoint"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck-web/pkg/transport"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/httputil"
	"github.com/MicroOps-cn/fuck/clients/tracing"
	"github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/log/flag"
	"github.com/MicroOps-cn/fuck/signals"
	w "github.com/MicroOps-cn/fuck/wrapper"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-logr/stdr"
	"github.com/oklog/oklog/pkg/group"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	cfgFile         string
	configDisplay   bool
	debugAddr       string
	httpExternalURL httputil.URL
	adminPrefix     string
	httpAddr        string
	openapiPath     string
	swaggerPath     string
	swaggerFilePath string
)

type Server struct {
	tracer          *sdktrace.TracerProvider
	svc             service.Service
	g               group.Group
	requestDuration metrics.Histogram
}

func (s *Server) Run() error {
	return s.g.Run()
}

func (s *Server) Add(execute func() error, interrupt func(error)) {
	s.g.Add(execute, interrupt)
}

func (s *Server) GetService() service.Service {
	return s.svc
}

func (s *Server) GetRequestDuration() metrics.Histogram {
	return s.requestDuration
}

func New(ctx context.Context, name string, logger kitlog.Logger, stopCh *signals.Handler) (server *Server, err error) {
	var s Server
	{
		if traceOptions := config.Get().Trace; traceOptions != nil {
			otel.SetLogger(stdr.New(stdlog.New(kitlog.NewStdlibAdapter(level.Info(logger)), "[restful]", stdlog.LstdFlags|stdlog.Lshortfile)))
			s.tracer, err = tracing.NewTraceProvider(ctx, config.Get().Trace)
			if err != nil {
				return nil, err
			}
			otel.SetTracerProvider(s.tracer)
			if http.DefaultClient.Transport == nil {
				http.DefaultClient.Transport = otelhttp.NewTransport(http.DefaultTransport)
			}
			tracing.SetTraceOptions(config.Get().Trace)
		}
	}

	// Create the (sparse) metrics we'll use in the service. They, too, are
	// dependencies that we pass to components that use them.

	{
		// Endpoint-level metrics.
		s.requestDuration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Name: "endpoint_invoke_duration_seconds",
			Help: "Tracks the latencies for Invoke endpoints.",
		}, []string{"method", "success"})
	}
	{
		httpLoginURL := httpExternalURL
		httpLoginURL.Path = path.Join(httpLoginURL.Path, adminPrefix, "account/login")
		ctx = context.WithValue(ctx, global.HTTPLoginURLKey, httpLoginURL.String())
		ctx = context.WithValue(ctx, global.HTTPExternalURLKey, httpExternalURL.String())
		ctx = context.WithValue(ctx, global.HTTPWebPrefixKey, adminPrefix)

		level.Info(logger).Log("msg", "Start service", "externalUrl", httpExternalURL, "adminPrefix", adminPrefix, "loginUrl", httpLoginURL)
		s.svc, err = service.New(ctx)
		if err != nil {
			return nil, err
		}
		if err = s.svc.LoadSystemConfig(ctx); err != nil {
			if strings.Contains(err.Error(), "no such table") {
				if err = s.svc.GetDB().Session(ctx).AutoMigrate(&models.SystemConfig{}); err != nil {
					panic(fmt.Errorf("failed to load system config: %s", err))
				}
			}
			if err = s.svc.LoadSystemConfig(ctx); err != nil {
				panic(fmt.Errorf("failed to load system config: %s", err))
			}
		}
	}

	{
		// The debug listener mounts the http.DefaultServeMux, and serves up
		// stuff like the Prometheus metrics route, the Go debug and profiling
		// routes, and so on.
		debugListener, err := net.Listen("tcp", debugAddr)
		if err != nil {
			level.Error(logger).Log("transport", "debug/HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
		s.g.Add(func() error {
			level.Info(logger).Log("msg", "Listening port", "transport", "debug/HTTP", "addr", debugAddr)
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			debugListener.Close()
			level.Debug(logger).Log("msg", "Listen closed", "transport", "debug/HTTP", "addr", debugAddr)
		})
	}
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			level.Error(logger).Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}

		s.g.Add(func() error {
			httpServer := http.NewServeMux()
			if len(swaggerPath) > 0 && len(openapiPath) > 0 && len(swaggerFilePath) > 0 {
				stat, err := os.Stat(swaggerFilePath)
				if err != nil {
					level.Error(logger).Log("err", err, "msg", "Failed to get swagger UI directory status, so disable that.")
				} else if stat.IsDir() {
					httpServer.Handle(swaggerPath, http.StripPrefix(swaggerPath, http.FileServer(http.Dir(swaggerFilePath))))
					level.Info(logger).Log("msg", fmt.Sprintf("enable Swagger UI on `%s` => %s", swaggerPath, swaggerFilePath))
				} else {
					level.Error(logger).Log("msg", " swagger UI local path is not directory, so disable that.")
				}
			}
			endpoints := endpoint.New(ctx, s.svc, s.requestDuration)
			httpHandler := transport.NewHTTPHandler(ctx, logger, endpoints, openapiPath)
			level.Info(logger).Log("msg", "Listening port", "transport", "HTTP", "addr", httpAddr)
			httpServer.Handle("/", httpHandler)
			serv := http.Server{Handler: httpServer, BaseContext: func(listener net.Listener) context.Context {
				return ctx
			}}

			return serv.Serve(httpListener)
		}, func(error) {
			httpListener.Close()
			level.Debug(logger).Log("msg", "Listen closed", "transport", "HTTP", "addr", httpAddr)
		})

		if _, ok := config.Get().Job.Scheduler.GetSchedulerBackend().(*config.JobOptions_Scheduler_XXLJob); ok {
			level.Debug(logger).Log("msg", "enable XXLJob service")
			jobGroupVersion := schema.GroupVersion{Group: "job", Version: "v1"}
			transport.RegisterServiceGenerator(transport.ServiceGeneratorFunc(transport.XXLJobService(jobGroupVersion)))
			ctx2, logger2 := log.NewContextLogger(ctx)
			_, port, _ := strings.Cut(httpAddr, ":")
			if len(port) == 0 {
				port = "80"
			}
			s.Add(func() error {
				jobRoot := fmt.Sprintf("%s/%s", transport.RootPath, jobGroupVersion.String())
				if err = s.svc.Register(ctx2, name, w.M(strconv.Atoi(port)), jobRoot); err != nil {
					level.Error(logger2).Log("service", "Job", "during", "Register", "err", err)
					return err
				}
				dur := time.NewTicker(time.Second * 30)
				for {
					select {
					case <-ctx.Done():
						return nil
					case <-dur.C:
						traceId := log.NewTraceId()
						newCtx, newLogger := log.NewContextLogger(ctx, log.WithTraceId(traceId), log.WithLogger(log.NewNopLogger()))
						err = s.svc.Register(newCtx, name, w.M(strconv.Atoi(port)), jobRoot)
						if err != nil {
							level.Error(newLogger).Log("service", "Job", "during", "Register", "err", err)
						}
					}
				}
			}, func(err error) {
				ctx2, logger2 = log.NewContextLogger(context.Background())
				if err = s.svc.Unregister(ctx2); err != nil {
					level.Error(logger2).Log("msg", "failed to unregister from job server", "err", err)
				} else {
					level.Info(logger2).Log("msg", "success to unregister from job server")
				}
			})
		}
	}
	{
		stopCh.Add(1)
		s.g.Add(func() error {
			stopCh.WaitRequest()
			return nil
		}, func(error) {
			if s.tracer != nil {
				timeoutCtx, closeCh := context.WithTimeout(context.Background(), time.Second*3)
				defer closeCh()
				if err = s.tracer.ForceFlush(timeoutCtx); err != nil {
					level.Debug(logger).Log("msg", "failed to force flush trace", "err", err)
					return
				}
				if err = s.tracer.Shutdown(timeoutCtx); err != nil {
					level.Debug(logger).Log("msg", "failed to force close trace", "err", err)
					return
				}
				stopCh.Done()
			}
		})
	}
	return &s, nil
}

func SetCommandFlags(rootCmd *cobra.Command) {
	cobra.OnInitialize(initParameter)
	// log level and format
	flag.AddFlags(rootCmd.PersistentFlags(), nil)
	config.AddFlags(rootCmd.Flags())
	global.WithAppConfigFlags(rootCmd.Flags())
	rootCmd.Flags().StringVar(&debugAddr, "debug.listen-address", ":8080", "Debug and metrics listen address")
	rootCmd.Flags().StringVar(&httpAddr, "http.listen-address", ":8081", "HTTP listen address")
	rootCmd.Flags().StringVar(&openapiPath, "http.openapi-path", "", "path of openapi")
	rootCmd.Flags().StringVar(&swaggerPath, "http.swagger-path", "/apidocs/", "path of swagger ui. If the value is empty, the swagger UI is disabled.")
	rootCmd.Flags().Var(&httpExternalURL, "http.external-url", "The URL under which web server is externally reachable (for example, if IDAS is served via a reverse proxy). Used for generating relative and absolute links back to web server itself. If the URL has a path portion, it will be used to prefix all HTTP endpoints served by web server. If omitted, relevant URL components will be derived automatically.")
	rootCmd.Flags().StringVar(&adminPrefix, "http.admin-prefix", "/admin/", "The path prefix of the static page. The default is the path of http.external-url.")
	rootCmd.Flags().StringVar(&swaggerFilePath, "swagger.file-path", "", "path of swagger ui local file. If the value is empty, the swagger UI is disabled.")
}

func initParameter() {
	logger := log.NewTraceLogger()

	if httpExternalURL.Scheme == "" {
		httpExternalURL.Scheme = "http"
	}
	if httpExternalURL.Path == "" {
		httpExternalURL.Path = "/"
	}
	if httpExternalURL.Path[len(httpExternalURL.Path)-1:] != "/" {
		httpExternalURL.Path = httpExternalURL.Path + "/"
	}

	if httpExternalURL.Host == "" {
		port := "80"
		if h, p, err := net.SplitHostPort(httpAddr); err == nil {
			port = p
			ip := net.ParseIP(h)
			if ip.IsLoopback() || ip.IsGlobalUnicast() {
				httpExternalURL.Host = httpAddr
			}
		}
		if httpExternalURL.Host == "" {
			httpExternalURL.Host = net.JoinHostPort("localhost", port)
			interfaces, err := net.Interfaces()
			if err != nil {
				level.Error(logger).Log("msg", "failed to get interface, please specify a valid http.external-url.")
			} else {
			loop:
				for _, iface := range interfaces {
					addrs, err := iface.Addrs()
					if err == nil {
						for _, addr := range addrs {
							ip, _, _ := net.ParseCIDR(addr.String())
							if ip.IsGlobalUnicast() {
								httpExternalURL.Host = net.JoinHostPort(ip.String(), port)
								break loop

							}
						}
					}
				}
			}
		}
	}
	if !strings.HasPrefix(adminPrefix, "/") {
		adminPrefix = "/" + adminPrefix
	}
	if !strings.HasSuffix(adminPrefix, "/") {
		adminPrefix = adminPrefix + "/"
	}
}

func AddCommand(rootCmd *cobra.Command) {
	SetCommandFlags(rootCmd)
	AddInitCommand(rootCmd)
	AddMigrateCommand(rootCmd)
	AddUserCommand(rootCmd)
	AddWeakPasswordCommand(rootCmd)
	AddSafeCommand(rootCmd)
	w.AddVersionFlags(func(shortVersion, fullVersion string) {
		rootCmd.Version = strings.TrimSpace(fullVersion)
	})
}
