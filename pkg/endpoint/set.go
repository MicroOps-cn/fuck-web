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

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(ctx context.Context, svc service.Service, duration metrics.Histogram) Set {
	logger := log.GetContextLogger(ctx)
	ps := Set{}.GetPermissionsDefine()

	var eps = sets.New[string]()
	injectEndpoint := func(name string, ep endpoint.Endpoint) endpoint.Endpoint {
		if eps.Has(name) {
			panic("duplicate endpoint: " + name)
		}
		psd := ps.Get(name)
		if len(psd) == 0 {
			panic("endpoint not found: " + name)
		} else if len(psd) > 1 {
			panic("duplicate endpoint define: " + name)
		}
		eps.Insert(name)
		if psd[0].RateLimit != nil {
			level.Debug(logger).Log("msg", "Injection rate limited to endpoints", "endpoint", name, "limit", psd[0].RateLimit)
			ep = ratelimit.NewErroringLimiter(psd[0].RateLimit)(ep)
		}
		ep = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(ep)
		ep = LoggingMiddleware(svc, name, ps)(ep)
		if duration != nil {
			ep = InstrumentingMiddleware(duration, name)(ep)
		}
		if psd[0].EnableAuth {
			ep = AuthorizationMiddleware(svc, name)(ep)
		}
		return ep
	}

	return Set{
		FileEndpoints: FileEndpoints{
			UploadFile:   injectEndpoint("UploadFile", MakeUploadFileEndpoint(svc)),
			DownloadFile: injectEndpoint("DownloadFile", MakeDownloadFileEndpoint(svc)),
		},
		UserEndpoints: UserEndpoints{
			CurrentUser:       injectEndpoint("CurrentUser", MakeCurrentUserEndpoint(svc)),
			ResetPassword:     injectEndpoint("ResetPassword", MakeResetUserPasswordEndpoint(svc)),
			ForgotPassword:    injectEndpoint("ForgotPassword", MakeForgotPasswordEndpoint(svc)),
			GetUsers:          injectEndpoint("GetUsers", MakeGetUsersEndpoint(svc)),
			DeleteUsers:       injectEndpoint("DeleteUsers", MakeDeleteUsersEndpoint(svc)),
			PatchUsers:        injectEndpoint("PatchUsers", MakePatchUsersEndpoint(svc)),
			UpdateUser:        injectEndpoint("UpdateUser", MakeUpdateUserEndpoint(svc)),
			GetUserInfo:       injectEndpoint("GetUserInfo", MakeGetUserInfoEndpoint(svc)),
			CreateUser:        injectEndpoint("CreateUser", MakeCreateUserEndpoint(svc)),
			PatchUser:         injectEndpoint("PatchUser", MakePatchUserEndpoint(svc)),
			DeleteUser:        injectEndpoint("DeleteUser", MakeDeleteUserEndpoint(svc)),
			CreateTOTPSecret:  injectEndpoint("CreateTOTPSecret", MakeCreateTOTPSecretEndpoint(svc)),
			CreateTOTP:        injectEndpoint("CreateTOTP", MakeCreateTOTPEndpoint(svc)),
			SendLoginCaptcha:  injectEndpoint("SendLoginCaptcha", MakeSendLoginCaptchaEndpoint(svc)),
			UpdateCurrentUser: injectEndpoint("UpdateCurrentUser", MakeUpdateCurrentUserEndpoint(svc)),
			PatchCurrentUser:  injectEndpoint("PatchCurrentUser", MakePatchCurrentUserEndpoint(svc)),
			ActivateAccount:   injectEndpoint("ActivateAccount", MakeActivateAccountEndpoint(svc)),
			SendActivateMail:  injectEndpoint("SendActivateMail", MakeSendActivationMailEndpoint(svc)),
		},
		SessionEndpoints: SessionEndpoints{
			GetSessions:              injectEndpoint("GetSessions", MakeGetSessionsEndpoint(svc)),
			GetCurrentUserSessions:   injectEndpoint("GetCurrentUserSessions", MakeGetCurrentUserSessionsEndpoint(svc)),
			DeleteCurrentUserSession: injectEndpoint("DeleteCurrentUserSession", MakeDeleteCurrentUserSessionEndpoint(svc)),
			DeleteSession:            injectEndpoint("DeleteSession", MakeDeleteSessionEndpoint(svc)),
			UserLogin:                injectEndpoint("UserLogin", MakeUserLoginEndpoint(svc)),
			UserOAuthLogin:           injectEndpoint("UserOAuthLogin", MakeUserOAuthLoginEndpoint(svc)),
			Authentication:           injectEndpoint("Authentication", MakeAuthenticationEndpoint(svc)),
			UserLogout:               injectEndpoint("UserLogout", MakeUserLogoutEndpoint(svc)),
			GetSessionByToken:        injectEndpoint("GetSessionByToken", MakeGetSessionByTokenEndpoint(svc)),
		},
		RoleEndpoints: RoleEndpoints{
			GetRoles:       injectEndpoint("GetRoles", MakeGetRolesEndpoint(svc)),
			DeleteRoles:    injectEndpoint("DeleteRoles", MakeDeleteRolesEndpoint(svc)),
			CreateRole:     injectEndpoint("CreateRole", MakeCreateRoleEndpoint(svc)),
			UpdateRole:     injectEndpoint("UpdateRole", MakeUpdateRoleEndpoint(svc)),
			DeleteRole:     injectEndpoint("DeleteRole", MakeDeleteRoleEndpoint(svc)),
			GetPermissions: injectEndpoint("GetPermissions", MakeGetPermissionsEndpoint(svc)),
		},
		ConfigEndpoints: ConfigEndpoints{
			GetSecurityConfig:   injectEndpoint("GetSecurityConfig", MakeGetSecurityConfigEndpoint(svc)),
			PatchSecurityConfig: injectEndpoint("PatchSecurityConfig", MakePatchSecurityConfigEndpoint(svc)),
		},
		EventEndpoints: EventEndpoints{
			GetEvents:               injectEndpoint("GetEvents", MakeGetEventsEndpoint(svc)),
			GetEventLogs:            injectEndpoint("GetEventLogs", MakeGetEventLogsEndpoint(svc)),
			GetCurrentUserEvents:    injectEndpoint("GetCurrentUserEvents", MakeGetCurrentUserEventsEndpoint(svc)),
			GetCurrentUserEventLogs: injectEndpoint("GetCurrentUserEventLogs", MakeGetCurrentUserEventLogsEndpoint(svc)),
		},
		GlobalEndpoints: GlobalEndpoints{
			GetGlobalConfig: injectEndpoint("GetGlobalConfig", MakeGetGlobalConfigEndpoint(svc)),
		},
	}
}

func (s Set) GetPermissionsDefine() models.Permissions {
	return GetPermissionsDefine(reflect.TypeOf(s))
}
