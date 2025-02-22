package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sync"

	jwtutils "github.com/MicroOps-cn/fuck/jwt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/prometheus/client_golang/prometheus"
	yaml "gopkg.in/yaml.v3"

	"github.com/MicroOps-cn/fuck-web/pkg/global"
)

var (
	configReloadSuccess = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: global.AppName,
		Name:      "config_last_reload_successful",
		Help:      "Blackbox exporter config loaded successfully.",
	})

	configReloadSeconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: global.AppName,
		Name:      "config_last_reload_success_timestamp_seconds",
		Help:      "Timestamp of the last successful configuration reload.",
	})
	safeCfg = newSafeConfig()
)

type onChange struct {
	f        func(o, n *Config) error
	fallback func(o, n *Config)
}

type safeConfig struct {
	sync.RWMutex
	C         *Config
	RC        *RuntimeConfig
	onChanges []onChange
}

func newSafeConfig() *safeConfig {
	return &safeConfig{
		C: &Config{},
		RC: &RuntimeConfig{
			Security: &RuntimeSecurityConfig{
				ForceEnableMfa:              false,
				PasswordComplexity:          0,
				PasswordMinLength:           0,
				PasswordExpireTime:          0,
				PasswordFailedLockThreshold: 0,
				PasswordFailedLockDuration:  0,
				PasswordHistory:             0,
				AccountInactiveLock:         0,
			},
		},
	}
}

func Get() *Config {
	return safeCfg.GetConfig()
}

func GetRuntimeConfig() *RuntimeConfig {
	return safeCfg.GetRuntimeConfig()
}

func SetRuntimeConfig(f func(c *RuntimeConfig)) {
	safeCfg.SetRuntimeConfig(f)
}

func (sc *safeConfig) SetRuntimeConfig(f func(c *RuntimeConfig)) {
	sc.Lock()
	defer sc.Unlock()
	f(sc.RC)
}

func (sc *safeConfig) GetRuntimeConfig() *RuntimeConfig {
	sc.RLock()
	defer sc.RUnlock()
	return sc.RC
}

func (sc *safeConfig) SetConfig(conf *Config) {
	sc.Lock()
	defer sc.Unlock()
	if sc.C != nil {
		for idx, f := range safeCfg.onChanges {
			if err := f.f(safeCfg.C, conf); err != nil {
				for i := 0; i <= idx; i++ {
					safeCfg.onChanges[i].fallback(safeCfg.GetConfig(), conf)
				}
			}
		}
	}
	sc.C = conf
	os.Setenv("APP_NAME", conf.GetAppName())
}

func (sc *safeConfig) GetConfig() *Config {
	sc.RLock()
	defer sc.RUnlock()
	return sc.C
}

type Converter struct {
	io.Reader
	name string
}

func (c *Converter) Name() string {
	return c.name
}

func (c *Converter) UnmarshalYAML(value *yaml.Node) error {
	vals := make(map[string]interface{})
	err := value.Decode(&vals)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(nil)
	c.Reader = buf
	return json.NewEncoder(buf).Encode(vals)
}

var _ yaml.Unmarshaler = &Converter{}

func NewConverter(name string, r io.Reader) *Converter {
	return &Converter{name: name, Reader: r}
}

func (sc *safeConfig) ReloadConfigFromYamlReader(logger log.Logger, yamlReader Reader) (err error) {
	defer func() {
		if err != nil {
			configReloadSuccess.Set(0)
		} else {
			configReloadSuccess.Set(1)
			configReloadSeconds.SetToCurrentTime()
		}
	}()
	cfgConvert := new(Converter)
	cfgConvert.name = yamlReader.Name()
	if err = yaml.NewDecoder(yamlReader).Decode(&cfgConvert); err != nil {
		return fmt.Errorf("error parse config file: %s", err)
	}
	return sc.ReloadConfigFromJSONReader(logger, cfgConvert)
}

type Reader interface {
	io.Reader
	Name() string
}

func (sc *safeConfig) ReloadConfigFromJSONReader(logger log.Logger, reader Reader) (err error) {
	defer func() {
		if err != nil {
			configReloadSuccess.Set(0)
		} else {
			configReloadSuccess.Set(1)
			configReloadSeconds.SetToCurrentTime()
		}
	}()

	c := Config{
		Global:   NewGlobalOptions(),
		Storage:  &Storages{},
		Security: &SecurityOptions{},
		Job: &JobOptions{
			Scheduler: &JobOptions_Scheduler{
				SchedulerBackend: &JobOptions_Scheduler_Local{
					Local: &JobOptions_LocalScheduler{},
				},
			},
		},
	}

	var unmarshaler jsonpb.Unmarshaler
	if err = unmarshaler.Unmarshal(reader, &c); err != nil {
		return fmt.Errorf("error unmarshal config: %s", err)
	} else if err = c.Init(); err != nil {
		return fmt.Errorf("error init config: %s", err)
	}
	if c.GetWorkspace() == nil {
		if absPath, err := filepath.Abs(path.Dir(reader.Name())); err != nil {
			c.SetWorkspace(path.Dir(reader.Name()))
			level.Debug(logger).Log("msg", "set workspace", "workspace", path.Dir(reader.Name()))
		} else {
			c.SetWorkspace(absPath)
			level.Debug(logger).Log("msg", "set workspace", "workspace", absPath)
		}
	}

	if c.Security == nil {
		c.Security = &SecurityOptions{}
	}
	if c.Security.Jwt == nil {
		if sc.C != nil && sc.C.Security != nil && sc.C.Security.Jwt != nil {
			c.Security.Jwt = sc.C.Security.Jwt
		} else {
			level.Warn(logger).Log("msg", "JWT key not set, will be automatically generated soon.")
			c.Security.Jwt, err = jwtutils.NewRandomRSAJWTConfig()
			if err != nil {
				return fmt.Errorf("failed to generate jwt config: %s", err)
			}
		}
	}
	sc.SetConfig(&c)
	return nil
}

func (sc *safeConfig) ReloadConfigFromFile(logger log.Logger, filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	ext := filepath.Ext(filename)
	if len(ext) > 1 {
		switch ext {
		case ".yml", ".yaml":
			return sc.ReloadConfigFromYamlReader(logger, r)
		case ".json":
			return sc.ReloadConfigFromJSONReader(logger, r)
		}
	}
	return nil
}

func ReloadConfigFromFile(logger log.Logger, filename string) error {
	return safeCfg.ReloadConfigFromFile(logger, filename)
}

func ReloadConfigFromYamlReader(logger log.Logger, r Reader) error {
	return safeCfg.ReloadConfigFromYamlReader(logger, r)
}

func SetConfig(cfg *Config) {
	safeCfg.SetConfig(cfg)
}
func OnConfigReload(f func(o, n *Config) error, fallback func(o, n *Config)) {
	safeCfg.onChanges = append(safeCfg.onChanges, onChange{f: f, fallback: fallback})
}
