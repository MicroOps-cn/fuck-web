package xxljob

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/MicroOps-cn/fuck/log"
	"github.com/go-kit/log/level"
	"github.com/spf13/afero"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Client struct {
	options    Options
	mux        sync.Mutex
	logDir     afero.Fs
	registered bool
	client     *http.Client
}

func (c *Client) OnJobHandlerChange(context.Context) {}

func (c *Client) GetConcurrency() int {
	return int(c.options.Concurrency)
}

func (c *Client) UnmarshalJSON(bytes []byte) error {
	c.options.Concurrency = int32(runtime.NumCPU())
	if err := json.Unmarshal(bytes, &c.options); err != nil {
		return err
	}

	return nil
}

type RegisterBody struct {
	RegistryGroup string `json:"registryGroup"`
	RegistryKey   string `json:"registryKey"`
	RegistryValue string `json:"registryValue"`
}

type ResponseBody struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func New(o Options) *Client {
	if o.Concurrency == 0 {
		o.Concurrency = int32(runtime.NumCPU())
	}
	dialer := net.Dialer{}
	transport := otelhttp.NewTransport(&http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}, otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
		if i := strings.LastIndex(r.URL.Path, "/"); i >= 0 {
			return "XXL-JOB." + r.URL.Path[i+1:]
		}
		return "XXL-JOB." + r.URL.Path
	}))

	return &Client{
		client:  &http.Client{Transport: transport},
		options: o,
	}
}

func (c *Client) GetLocalAddress(ctx context.Context, port int, root string) (string, error) {
	localAddress := c.options.LocalAddress
	if len(localAddress) != 0 {
		return localAddress, nil
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		var ip *net.IP
		switch ad := addr.(type) {
		case *net.IPNet:
			ip = &ad.IP
		case *net.IPAddr:
			ip = &ad.IP
		default:
			continue
		}
		if ip != nil {
			if ip.IsLoopback() || ip.IsUnspecified() || ip.IsLinkLocalUnicast() {
				continue
			}
			root = strings.TrimPrefix(root, "/")
			return fmt.Sprintf("http://%s:%d/%s", ip.String(), port, root), nil
		}
	}
	return "", nil
}

func (c *Client) Register(ctx context.Context, name string, port int, root string) error {
	localAddress, err := c.GetLocalAddress(ctx, port, root)
	if err != nil {
		return fmt.Errorf("failed to get local address: %s", err)
	} else if len(c.options.AppName) != 0 {
		name = c.options.AppName
	}
	if !c.registered {
		logger := log.GetContextLogger(ctx)
		level.Debug(logger).Log("msg", "register to job server", "name", name, "localAddress", localAddress)
		c.options.LocalAddress = localAddress
		if len(c.options.LogPath) == 0 {
			c.options.LogPath = "logs/job"
		}
		_, err = os.Stat(c.options.LogPath)
		if err != nil && os.IsNotExist(err) {
			if err = os.MkdirAll(c.options.LogPath, 0o755); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
		c.logDir = afero.NewBasePathFs(afero.NewOsFs(), c.options.LogPath)
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	err = c.Do(ctx, "POST", "/api/registry", RegisterBody{
		RegistryGroup: "EXECUTOR",
		RegistryKey:   name,
		RegistryValue: localAddress,
	})
	if err != nil {
		return err
	}
	c.options.AppName = name
	c.registered = true
	return nil
}

func (c *Client) Unregister(ctx context.Context) error {
	logger := log.GetContextLogger(ctx)
	if len(c.options.LocalAddress) == 0 {
		return fmt.Errorf("unknown local address")
	}
	level.Debug(logger).Log("msg", "unregister to job server", "name", c.options.AppName, "localAddress", c.options.LocalAddress)
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.Do(ctx, "POST", "/api/registryRemove", RegisterBody{
		RegistryGroup: "EXECUTOR",
		RegistryKey:   c.options.AppName,
		RegistryValue: c.options.LocalAddress,
	})
}

type CallbackBody struct {
	LogId      int32       `json:"logId"`
	LogDateTim int64       `json:"logDateTim"`
	HandleCode int         `json:"handleCode"`
	HandleMsg  interface{} `json:"handleMsg"`
}

func (c *Client) Callback(ctx context.Context, logId int32, code int, msg string) error {
	if len(c.options.LocalAddress) == 0 {
		return fmt.Errorf("unknown local address")
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.Do(ctx, "POST", "/api/callback", []CallbackBody{{
		LogId:      logId,
		LogDateTim: time.Now().Unix(),
		HandleCode: code,
		HandleMsg:  msg,
	}})
}

func (c *Client) Do(ctx context.Context, method string, api string, body interface{}) (err error) {
	logger := log.GetContextLogger(ctx)
	level.Debug(logger).Log("msg", "request job server", "method", method, "api", api)
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*30)
	defer cancelFunc()
	addr := c.options.ServerAddress
	u, err := url.Parse(addr)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, api)
	var req *http.Request
	switch b := body.(type) {
	case io.Reader:
		if req, err = http.NewRequest(method, u.String(), b); err != nil {
			return fmt.Errorf("failed to create request: %s", err)
		}
	case string:
		if req, err = http.NewRequest(method, u.String(), bytes.NewBufferString(b)); err != nil {
			return fmt.Errorf("failed to create request: %s", err)
		}
	default:
		rawData, err := json.Marshal(body)
		if err != nil {
			return err
		}
		if req, err = http.NewRequest(method, u.String(), bytes.NewBufferString(string(rawData))); err != nil {
			return fmt.Errorf("failed to create request: %s", err)
		}
		req.Header.Set("Content-Type", "application/json")
	}
	if len(c.options.Token) != 0 {
		req.Header.Set("XXL-JOB-ACCESS-TOKEN", c.options.Token)
	}

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("failed to request job server: %s", err)
	}
	defer resp.Body.Close()
	var r ResponseBody
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return fmt.Errorf("failed to decode job server response: %s", err)
	}
	if r.Code != 200 {
		return fmt.Errorf("failed to request job server: code=%d,msg=%s", r.Code, r.Msg)
	}
	return err
}
