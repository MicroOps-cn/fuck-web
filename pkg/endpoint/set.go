package endpoint

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/sets"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/log/level"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
)

type XXLJobEndpoints struct {
	ExecutorBeat     endpoint.Endpoint ` auth:"false" audit:"false"`
	ExecutorIdleBeat endpoint.Endpoint ` auth:"false" audit:"false"`
	ExecutorKill     endpoint.Endpoint ` auth:"false" audit:"false"`
	ExecutorLog      endpoint.Endpoint ` auth:"false" audit:"false"`
	ExecutorRunJob   endpoint.Endpoint ` auth:"false" audit:"false"`
}

type UserEndpoints struct {
	GetUsers          endpoint.Endpoint `description:"Get user list" role:"admin|viewer" audit:"false"`
	DeleteUsers       endpoint.Endpoint `description:"Batch delete users" role:"admin" audit:"false"`
	PatchUsers        endpoint.Endpoint `description:"Batch modify user information (incremental)" role:"admin" audit:"true"`
	UpdateUser        endpoint.Endpoint `description:"Modify user information" role:"admin" audit:"true"`
	GetUserInfo       endpoint.Endpoint `description:"Get user details" role:"admin|viewer" audit:"false"`
	CreateUser        endpoint.Endpoint `description:"Create a user" role:"admin" audit:"true"`
	PatchUser         endpoint.Endpoint `description:"Modify user information (incremental)" role:"admin" audit:"true"`
	DeleteUser        endpoint.Endpoint `description:"Delete a user" role:"admin" audit:"true"`
	ForgotPassword    endpoint.Endpoint `auth:"false" audit:"true"`
	ResetPassword     endpoint.Endpoint `auth:"false" audit:"true"`
	CurrentUser       endpoint.Endpoint `auth:"false" audit:"false"`
	CreateTOTPSecret  endpoint.Endpoint `auth:"false" audit:"false"`
	CreateTOTP        endpoint.Endpoint `auth:"false" audit:"true"`
	UnbindTOTP        endpoint.Endpoint `auth:"false" audit:"true"`
	SendLoginCaptcha  endpoint.Endpoint `auth:"false" audit:"true"`
	UpdateCurrentUser endpoint.Endpoint `auth:"false" audit:"true"`
	PatchCurrentUser  endpoint.Endpoint `auth:"false" audit:"true"`
	SendActivateMail  endpoint.Endpoint `description:"Send activation link to user mail" role:"admin" audit:"true"`
	ActivateAccount   endpoint.Endpoint `auth:"false" audit:"true"`
}

type SessionEndpoints struct {
	GetSessions              endpoint.Endpoint `description:"Get the user's session list" role:"admin" audit:"false"`
	DeleteSession            endpoint.Endpoint `description:"Delete the user's session" role:"admin" audit:"true"`
	GetCurrentUserSessions   endpoint.Endpoint `description:"Get current user session list" auth:"false" audit:"false"`
	DeleteCurrentUserSession endpoint.Endpoint `description:"Get current user session list" auth:"false" audit:"false"`
	UserLogin                endpoint.Endpoint `auth:"false" audit:"true"`
	UserOAuthLogin           endpoint.Endpoint `auth:"false" audit:"false"`
	UserLogout               endpoint.Endpoint `auth:"false" audit:"true"`
	GetSessionByToken        endpoint.Endpoint `auth:"false" audit:"false"`
	Authentication           endpoint.Endpoint `auth:"false" audit:"false"`
	SessionRenewal           endpoint.Endpoint `auth:"false" audit:"false"`
}

type RoleEndpoints struct {
	GetPermissions endpoint.Endpoint `description:"Get permission list" role:"admin|viewer" audit:"false"`
	GetRoles       endpoint.Endpoint `description:"Get role list" role:"admin|viewer" audit:"false"`
	DeleteRoles    endpoint.Endpoint `description:"Batch delete roles" role:"admin"`
	CreateRole     endpoint.Endpoint `description:"Create a role" role:"admin"`
	UpdateRole     endpoint.Endpoint `description:"Modify role information" role:"admin"`
	DeleteRole     endpoint.Endpoint `description:"Delete a role" role:"admin"`
}

type PageEndpoints struct {
	GetPages   endpoint.Endpoint `description:"Get page list" role:"admin|viewer" audit:"false"`
	GetPage    endpoint.Endpoint `description:"Get page" role:"admin|viewer" audit:"false"`
	CreatePage endpoint.Endpoint `description:"Create a page" role:"admin"`
	UpdatePage endpoint.Endpoint `description:"Modify page information" role:"admin"`
	DeletePage endpoint.Endpoint `description:"Delete a page" role:"admin"`
	PatchPages endpoint.Endpoint `description:"Patch pages" role:"admin"`

	GetPageDatas   endpoint.Endpoint `description:"Get data list of page" role:"admin|viewer" audit:"false"`
	GetPageData    endpoint.Endpoint `description:"Get a data of page" role:"admin|viewer" audit:"false"`
	CreatePageData endpoint.Endpoint `description:"Create a data of page" role:"admin"`
	UpdatePageData endpoint.Endpoint `description:"Modify a data of page" role:"admin"`
	DeletePageData endpoint.Endpoint `description:"Delete a data of page" role:"admin"`
	PatchPageDatas endpoint.Endpoint `description:"Patch page data" role:"admin"`
}

type FileEndpoints struct {
	UploadFile   endpoint.Endpoint `name:"" description:"Upload files to the server" auth:"false" audit:"false"`
	DownloadFile endpoint.Endpoint `auth:"false" audit:"false"`
}

type ProxyEndpoints struct {
	ProxyRequest   endpoint.Endpoint `auth:"false" audit:"false"`
	GetProxyConfig endpoint.Endpoint `auth:"false" audit:"false"`
}

type ConfigEndpoints struct {
	GetSecurityConfig   endpoint.Endpoint `description:"Get security config." role:"admin" audit:"false"`
	PatchSecurityConfig endpoint.Endpoint `description:"Patch security config." role:"admin"`
}

type GlobalEndpoints struct {
	GetGlobalConfig endpoint.Endpoint `description:"Get login type." auth:"false" audit:"false"`
}

type EventEndpoints struct {
	GetEvents               endpoint.Endpoint `description:"Get events." role:"admin" audit:"false"`
	GetEventLogs            endpoint.Endpoint `description:"Get event logs." role:"admin" audit:"false"`
	GetCurrentUserEvents    endpoint.Endpoint `description:"Get current user events." auth:"false" audit:"false"`
	GetCurrentUserEventLogs endpoint.Endpoint `description:"Get current user event logs." auth:"false" audit:"false"`
}

// Set collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	UserEndpoints    `name:"User" description:"User management"`
	SessionEndpoints `name:"Session" description:"User session management"`
	RoleEndpoints    `name:"Role" description:"Role of current platform"`
	FileEndpoints    `name:"File" description:"File"`
	ConfigEndpoints  `name:"Config" description:"System Config Manage"`
	EventEndpoints   `name:"Event" description:"Event"`
	GlobalEndpoints  `name:"Global" description:"Global"`
	XXLJobEndpoints  `name:"XXLJob" description:"XXlJob"`
}

func GetPermissionsDefine(typeOf reflect.Type) models.Permissions {
	var ret models.Permissions
	for typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}
	for i := 0; i < typeOf.NumField(); i++ {
		var p models.Permission
		field := typeOf.Field(i)
		if p.Name = field.Tag.Get("name"); len(p.Name) == 0 {
			p.Name = field.Name
		}

		p.Description = field.Tag.Get("description")
		if p.RateLimit = config.Get().GetRateLimit(p.Name); p.RateLimit == nil {
			rl := field.Tag.Get("ratelimit")
			switch rl {
			case "0", "false":
			case "":
				if p.RateLimit = config.Get().GetRateLimit(""); p.RateLimit == nil {
					p.RateLimit = config.NewLimiterWrapper(rate.NewLimiter(rate.Limit(10), 100))
				}
			default:
				rateLimit, err := strconv.ParseFloat(rl, 64)
				if err != nil {
					panic(fmt.Errorf("failed to parse rate limit config of endpoint %s ", p.Name))
				}
				p.RateLimit = config.NewLimiterWrapper(rate.NewLimiter(rate.Limit(rateLimit), 100))
			}
		}
		if field.Type.Kind() == reflect.Struct {
			if auth := field.Tag.Get("auth"); len(auth) == 0 || auth == "true" {
				p.EnableAuth = true
			}
			p.Children = GetPermissionsDefine(field.Type)
			if len(p.Children) > 0 {
				ret = append(ret, &p)
				continue
			}
		} else if field.Type.Kind() == reflect.Func {
			if audit := field.Tag.Get("audit"); len(audit) == 0 || audit == "true" {
				p.EnableAudit = true
			}

			if auth := field.Tag.Get("auth"); len(auth) == 0 || auth == "true" {
				p.EnableAuth = true
			}
			if p.EnableAuth {
				if role := field.Tag.Get("role"); len(role) > 0 {
					p.Role = strings.Split(role, "|")
				}
			}
			ret = append(ret, &p)
		}
	}
	return ret
}

var (
	permissions models.Permissions = nil
	endpointSet                    = sets.New[string]()
)

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(ctx context.Context, svc service.Service, duration metrics.Histogram) Set {
	injectStdEndpoint := func(name string, ep endpoint.Endpoint) endpoint.Endpoint {
		return WithMiddleware(ctx, name, svc, duration, ep)
	}
	injectEndpoint := func(name string, ep Endpoint) endpoint.Endpoint {
		return WithMiddleware(ctx, name, svc, duration, func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return ep(ctx, request)
		})
	}

	return Set{
		FileEndpoints: FileEndpoints{
			UploadFile:   injectStdEndpoint("UploadFile", MakeUploadFileEndpoint(svc)),
			DownloadFile: injectStdEndpoint("DownloadFile", MakeDownloadFileEndpoint(svc)),
		},
		UserEndpoints: UserEndpoints{
			CurrentUser:       injectStdEndpoint("CurrentUser", MakeCurrentUserEndpoint(svc)),
			ResetPassword:     injectStdEndpoint("ResetPassword", MakeResetUserPasswordEndpoint(svc)),
			ForgotPassword:    injectStdEndpoint("ForgotPassword", MakeForgotPasswordEndpoint(svc)),
			GetUsers:          injectStdEndpoint("GetUsers", MakeGetUsersEndpoint(svc)),
			DeleteUsers:       injectStdEndpoint("DeleteUsers", MakeDeleteUsersEndpoint(svc)),
			PatchUsers:        injectStdEndpoint("PatchUsers", MakePatchUsersEndpoint(svc)),
			UpdateUser:        injectStdEndpoint("UpdateUser", MakeUpdateUserEndpoint(svc)),
			GetUserInfo:       injectStdEndpoint("GetUserInfo", MakeGetUserInfoEndpoint(svc)),
			CreateUser:        injectStdEndpoint("CreateUser", MakeCreateUserEndpoint(svc)),
			PatchUser:         injectStdEndpoint("PatchUser", MakePatchUserEndpoint(svc)),
			DeleteUser:        injectStdEndpoint("DeleteUser", MakeDeleteUserEndpoint(svc)),
			CreateTOTPSecret:  injectStdEndpoint("CreateTOTPSecret", MakeCreateTOTPSecretEndpoint(svc)),
			CreateTOTP:        injectStdEndpoint("CreateTOTP", MakeCreateTOTPEndpoint(svc)),
			SendLoginCaptcha:  injectStdEndpoint("SendLoginCaptcha", MakeSendLoginCaptchaEndpoint(svc)),
			UpdateCurrentUser: injectStdEndpoint("UpdateCurrentUser", MakeUpdateCurrentUserEndpoint(svc)),
			PatchCurrentUser:  injectStdEndpoint("PatchCurrentUser", MakePatchCurrentUserEndpoint(svc)),
			ActivateAccount:   injectStdEndpoint("ActivateAccount", MakeActivateAccountEndpoint(svc)),
			SendActivateMail:  injectStdEndpoint("SendActivateMail", MakeSendActivationMailEndpoint(svc)),
		},
		SessionEndpoints: SessionEndpoints{
			GetSessions:              injectStdEndpoint("GetSessions", MakeGetSessionsEndpoint(svc)),
			GetCurrentUserSessions:   injectStdEndpoint("GetCurrentUserSessions", MakeGetCurrentUserSessionsEndpoint(svc)),
			DeleteCurrentUserSession: injectStdEndpoint("DeleteCurrentUserSession", MakeDeleteCurrentUserSessionEndpoint(svc)),
			DeleteSession:            injectStdEndpoint("DeleteSession", MakeDeleteSessionEndpoint(svc)),
			UserLogin:                injectStdEndpoint("UserLogin", MakeUserLoginEndpoint(svc)),
			UserOAuthLogin:           injectStdEndpoint("UserOAuthLogin", MakeUserOAuthLoginEndpoint(svc)),
			Authentication:           injectStdEndpoint("Authentication", MakeAuthenticationEndpoint(svc)),
			UserLogout:               injectStdEndpoint("UserLogout", MakeUserLogoutEndpoint(svc)),
			GetSessionByToken:        injectStdEndpoint("GetSessionByToken", MakeGetSessionByTokenEndpoint(svc)),
		},
		RoleEndpoints: RoleEndpoints{
			GetRoles:       injectStdEndpoint("GetRoles", MakeGetRolesEndpoint(svc)),
			DeleteRoles:    injectStdEndpoint("DeleteRoles", MakeDeleteRolesEndpoint(svc)),
			CreateRole:     injectStdEndpoint("CreateRole", MakeCreateRoleEndpoint(svc)),
			UpdateRole:     injectStdEndpoint("UpdateRole", MakeUpdateRoleEndpoint(svc)),
			DeleteRole:     injectStdEndpoint("DeleteRole", MakeDeleteRoleEndpoint(svc)),
			GetPermissions: injectStdEndpoint("GetPermissions", MakeGetPermissionsEndpoint(svc)),
		},
		ConfigEndpoints: ConfigEndpoints{
			GetSecurityConfig:   injectStdEndpoint("GetSecurityConfig", MakeGetSecurityConfigEndpoint(svc)),
			PatchSecurityConfig: injectStdEndpoint("PatchSecurityConfig", MakePatchSecurityConfigEndpoint(svc)),
		},
		EventEndpoints: EventEndpoints{
			GetEvents:               injectStdEndpoint("GetEvents", MakeGetEventsEndpoint(svc)),
			GetEventLogs:            injectStdEndpoint("GetEventLogs", MakeGetEventLogsEndpoint(svc)),
			GetCurrentUserEvents:    injectStdEndpoint("GetCurrentUserEvents", MakeGetCurrentUserEventsEndpoint(svc)),
			GetCurrentUserEventLogs: injectStdEndpoint("GetCurrentUserEventLogs", MakeGetCurrentUserEventLogsEndpoint(svc)),
		},
		GlobalEndpoints: GlobalEndpoints{
			GetGlobalConfig: injectStdEndpoint("GetGlobalConfig", MakeGetGlobalConfigEndpoint(svc)),
		},
		XXLJobEndpoints: XXLJobEndpoints{
			ExecutorBeat:     injectEndpoint("ExecutorBeat", MakeExecutorBeatEndpoint(svc)),
			ExecutorIdleBeat: injectEndpoint("ExecutorIdleBeat", MakeExecutorIdleBeatEndpoint(svc)),
			ExecutorKill:     injectEndpoint("ExecutorKill", MakeExecutorKillEndpoint(svc)),
			ExecutorLog:      injectEndpoint("ExecutorLog", MakeExecutorLogEndpoint(svc)),
			ExecutorRunJob:   injectEndpoint("ExecutorRunJob", MakeExecutorRunJobEndpoint(svc)),
		},
	}
}

func (s Set) GetPermissionsDefine() models.Permissions {
	return GetPermissionsDefine(reflect.TypeOf(s))
}

func GetPermissions() models.Permissions {
	if permissions == nil {
		for _, set := range registeredEndpointSet {
			permissions = append(permissions, GetPermissionsDefine(reflect.TypeOf(set))...)
		}
	}
	return permissions
}

var registeredEndpointSet []interface{}

func RegisterEndpointSet(set ...interface{}) {
	registeredEndpointSet = append(registeredEndpointSet, set...)
}

func init() {
	RegisterEndpointSet(&Set{})
}
func WithMiddleware(ctx context.Context, name string, svc service.Service, dur metrics.Histogram, ep endpoint.Endpoint) endpoint.Endpoint {

	logger := log.GetContextLogger(ctx)
	if permissions == nil {
		for _, set := range registeredEndpointSet {
			permissions = append(permissions, GetPermissionsDefine(reflect.TypeOf(set))...)
		}
	}
	if endpointSet.Has(name) {
		panic("duplicate endpoint: " + name)
	}
	psd := permissions.Get(name)
	if len(psd) == 0 {
		panic("endpoint not found: " + name)
	} else if len(psd) > 1 {
		panic("duplicate endpoint define: " + name)
	}
	endpointSet.Insert(name)
	if psd[0].RateLimit != nil {
		level.Debug(logger).Log("msg", "Injection rate limited to endpoints", "endpoint", name, "limit", psd[0].RateLimit)
		ep = ratelimit.NewErroringLimiter(psd[0].RateLimit)(ep)
	}
	ep = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(ep)
	ep = LoggingMiddleware(svc, name, permissions)(ep)
	if dur != nil {
		ep = InstrumentingMiddleware(dur, name)(ep)
	}
	if psd[0].EnableAuth {
		ep = AuthorizationMiddleware(svc, name)(ep)
	}
	return ep
}
