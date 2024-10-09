package global

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/stoewer/go-strcase"
)

var (
	AppName         = "fuck-web"
	LoginSession    = ""
	CookieAutoLogin = ""
	RedisKeyPrefix  = ""
)

type appNameHandler struct{}

func (a appNameHandler) String() string {
	return AppName
}

func (a appNameHandler) Set(s string) error {
	lowerAppName := strcase.SnakeCase(strings.NewReplacer(" ", "_", "-", "_").Replace(s))
	AppName = strings.ReplaceAll(lowerAppName, "_", "-")
	LoginSession = fmt.Sprintf("%s_LOGIN_SESSION", strings.ToUpper(lowerAppName))

	CookieAutoLogin = fmt.Sprintf("%s_AUTO_LOGIN", strings.ToUpper(lowerAppName))
	RedisKeyPrefix = strings.ToUpper(lowerAppName)
	return nil
}

func (a appNameHandler) Type() string {
	return "string"
}

type appSessionNameHandler struct{}

func (appSessionNameHandler) String() string {
	if LoginSession == "" {
		return "<UpperSnakeCase(app.name)>_LOGIN_SESSION"
	}

	lowerAppName := strcase.SnakeCase(strings.NewReplacer(" ", "_", "-", "_").Replace(AppName))
	if LoginSession == fmt.Sprintf("%s_LOGIN_SESSION", strings.ToUpper(lowerAppName)) {
		return "<UpperSnakeCase(app.name)>_LOGIN_SESSION"
	}
	return LoginSession
}

func (a appSessionNameHandler) Set(s string) error {
	LoginSession = s
	return nil
}

func (a appSessionNameHandler) Type() string {
	return "string"
}

func WithAppConfigFlags(flags *pflag.FlagSet) {
	flags.Var(&appNameHandler{}, "app.name", "application name.")
	flags.Var(&appSessionNameHandler{}, "app.session-name", "The session/cookie name used during login.")
}

func init() {
	err := (&appNameHandler{}).Set(AppName)
	if err != nil {
		fmt.Println("invalid app name", err)
		os.Exit(1)
	}
}

const (
	RestfulRequestContextName   = "__restful_request__"
	RestfulResponseContextName  = "__restful_response__"
	MetaUser                    = "__user__"
	MetaNeedLogin               = "__need_login__"
	MetaUpdateLastSeen          = "__need_login__"
	MetaForceOk                 = "__force_ok__"
	MetaSensitiveData           = "__sensitive_data__"
	MetaAutoRedirectToLoginPage = "__auto_redirect_to_login_page__"
	ActiveExpiration            = 7 * 24 * time.Hour
	AuthCodeExpiration          = 5 * time.Minute
	TokenExpiration             = 1 * time.Hour
	RefreshTokenExpiration      = 30 * 24 * time.Hour
	ResetPasswordExpiration     = 30 * time.Minute
	LoginSessionExpiresFormat   = "Mon, 02-Jan-06 15:04:05 MST"
	HTTPExternalURLKey          = "__http_external_url__"
	HTTPLoginURLKey             = "__http_login_url__"
	HTTPWebPrefixKey            = "__http_web_prefix__"
)
