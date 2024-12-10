package transport

import (
	"context"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-openapi/spec"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/MicroOps-cn/fuck-web/pkg/endpoint"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
)

var apiServiceSet = []func(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService){
	UserService,
	FileService,
	SessionService,
	CurrentUserService,
	PermissionService,
	RoleService,
	ConfigService,
	EventService,
	GlobalService,
}

var serviceGeneratorSet []ServiceGenerator

type ServiceGeneratorFunc func(context.Context, []httptransport.ServerOption, endpoint.Set) (spec.Tag, []*restful.WebService)

func (w ServiceGeneratorFunc) WebServices(ctx context.Context, options []httptransport.ServerOption, set endpoint.Set) (spec.Tag, []*restful.WebService) {
	return w(ctx, options, set)
}

type ServiceGenerator interface {
	WebServices(context.Context, []httptransport.ServerOption, endpoint.Set) (spec.Tag, []*restful.WebService)
}

func RegisterServiceGenerator(g ...ServiceGenerator) {
	serviceGeneratorSet = append(serviceGeneratorSet, g...)
}

// UserService User Manager Service for restful Http container
func UserService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "users", Description: "Managing users"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.GET("").
		To(NewKitHTTPServer[endpoint.GetUsersRequest](ctx, endpoints.GetUsers, options)).
		Operation("getUsers").
		Doc("Get user list.").
		Params(StructToQueryParams(endpoint.GetUsersRequest{})...).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetUsersResponse{}),
	)
	v1ws.Route(v1ws.PATCH("").
		To(NewKitHTTPServer[endpoint.PatchUsersRequest](ctx, endpoints.PatchUsers, options)).
		Operation("patchUsers").
		Reads(endpoint.PatchUsersRequest{}).
		Doc("Batch update user information(Incremental).").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseTotalResponse{}),
	)
	v1ws.Route(v1ws.DELETE("").
		To(NewKitHTTPServer[endpoint.DeleteUsersRequest](ctx, endpoints.DeleteUsers, options)).
		Operation("deleteUsers").
		Doc("Delete users in batch.").
		Reads(endpoint.DeleteUsersRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseTotalResponse{}),
	)
	v1ws.Route(v1ws.POST("").
		To(NewKitHTTPServer[endpoint.CreateUserRequest](ctx, endpoints.CreateUser, options)).
		Operation("createUser").
		Doc("Create a user.").
		Reads(endpoint.CreateUserRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)
	v1ws.Route(v1ws.GET("/{id}").
		To(NewKitHTTPServer[endpoint.GetUserRequest](ctx, endpoints.GetUserInfo, options)).
		Operation("getUserInfo").
		Param(v1ws.PathParameter("id", "identifier of the user").DataType("string")).
		Doc("Get user information.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetUserResponse{}),
	)
	v1ws.Route(v1ws.PUT("/{id}").
		To(NewKitHTTPServer[endpoint.UpdateUserRequest](ctx, endpoints.UpdateUser, options)).
		Operation("updateUser").
		Param(v1ws.PathParameter("id", "identifier of the user").DataType("string")).
		Doc("Update user information(full).").
		Reads(endpoint.UpdateUserRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.UpdateUserRequest{}),
	)
	v1ws.Route(v1ws.PATCH("/{id}").
		To(NewKitHTTPServer[endpoint.PatchUserRequest](ctx, endpoints.PatchUser, options)).
		Operation("patchUser").
		Param(v1ws.PathParameter("id", "identifier of the user").DataType("string")).
		Doc("Update user information(Incremental).").
		Reads(endpoint.PatchUserRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.PatchUserResponse{}),
	)

	v1ws.Route(v1ws.DELETE("/{id}").
		To(NewKitHTTPServer[endpoint.DeleteUserRequest](ctx, endpoints.DeleteUser, options)).
		Operation("deleteUser").
		Param(v1ws.PathParameter("id", "identifier of the user").DataType("string")).
		Doc("Delete user.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.POST("/sendActivateMail").
		To(NewKitHTTPServer[endpoint.SendActivationMailRequest](ctx, endpoints.SendActivateMail, options)).
		Operation("sendActivateMail").
		Reads(endpoint.SendActivationMailRequest{}).
		Doc("Send account activation email.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	return tag, []*restful.WebService{v1ws}
}

func FileService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "files", Description: "Managing files"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.POST("").
		To(NewKitHTTPServer[struct{}](ctx, endpoints.UploadFile, options)).
		Operation("uploadFile").
		Consumes("multipart/form-data").
		Doc("Upload file").
		Param(v1ws.MultiPartFormParameter("files", "files").AllowMultiple(true).DataType("file")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(global.MetaSensitiveData, true).
		Returns(200, "OK", endpoint.FileUploadResponse{}),
	)
	v1ws.Route(v1ws.GET("/{id}").
		To(NewSimpleKitHTTPServer[endpoint.FileDownloadRequest](ctx, endpoints.DownloadFile, DecodeHTTPRequest[endpoint.FileDownloadRequest], NoopEncodeHTTPResponse, options)).
		Operation("downloadFile").
		Param(v1ws.PathParameter("id", "identifier of the file").DataType("string").Required(true)).
		Doc("Download/View File").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)
	return tag, []*restful.WebService{v1ws}
}

func SessionService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "sessions", Description: "Managing sessions"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.GET("").
		To(NewKitHTTPServer[endpoint.GetSessionsRequest](ctx, endpoints.GetSessions, options)).
		Operation("getSessions").
		Doc("Get session list.").
		Params(StructToQueryParams(endpoint.GetSessionsRequest{})...).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetSessionsResponse{}),
	)
	v1ws.Route(v1ws.DELETE("/{id}").
		To(NewKitHTTPServer[endpoint.DeleteSessionRequest](ctx, endpoints.DeleteSession, options)).
		Operation("deleteSession").
		Param(v1ws.PathParameter("id", "identifier of the session").DataType("string").Required(true)).
		Doc("Expire a session.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)
	return tag, []*restful.WebService{v1ws}
}

func CurrentUserService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "user", Description: "Current user service"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.POST("/login").
		To(NewKitHTTPServer[endpoint.UserLoginRequest](ctx, endpoints.UserLogin, options)).
		Operation("userLogin").
		Doc("User login.").
		Reads(endpoint.UserLoginRequest{}).
		Metadata(global.MetaNeedLogin, false).
		Metadata(global.MetaSensitiveData, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.UserLoginResponse{}),
	)
	v1ws.Route(v1ws.GET("/oauth/{id}").
		To(NewKitHTTPServer[endpoint.OAuthLoginRequest](ctx, endpoints.UserOAuthLogin, options)).
		Operation("userOAuthLogin").
		Doc("OAuth login.").
		Param(v1ws.PathParameter("id", "identifier of the oauth").DataType("string").Required(true)).
		Metadata(global.MetaNeedLogin, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(302, "OK", endpoint.UserLoginResponse{}),
	)
	v1ws.Route(v1ws.POST("/logout").
		To(NewKitHTTPServer[struct{}](ctx, endpoints.UserLogout, options)).
		Operation("userLogout").
		Doc("User logout.").
		Consumes("*/*").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.GET("").
		To(NewKitHTTPServer[struct{}](ctx, endpoints.CurrentUser, options)).
		Operation("currentUser").
		Doc("Get current login user information.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(global.MetaUpdateLastSeen, true).
		Returns(200, "OK", endpoint.GetUserResponse{}),
	)

	v1ws.Route(v1ws.GET("events").
		To(NewKitHTTPServer[endpoint.GetCurrentUserEventsRequest](ctx, endpoints.GetCurrentUserEvents, options)).
		Operation("currentUserEvents").
		Params(StructToQueryParams(endpoint.GetCurrentUserEventsRequest{})...).
		Doc("Get current login user's events.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetCurrentUserEventsResponse{}),
	)

	v1ws.Route(v1ws.GET("eventLogs").
		To(NewKitHTTPServer[endpoint.GetCurrentUserEventLogsRequest](ctx, endpoints.GetCurrentUserEventLogs, options)).
		Operation("currentUserEventLogs").
		Params(StructToQueryParams(endpoint.GetCurrentUserEventLogsRequest{})...).
		Doc("Get current login user's event logs.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetCurrentUserEventLogsResponse{}),
	)

	v1ws.Route(v1ws.PUT("").
		To(NewKitHTTPServer[endpoint.UpdateUserRequest](ctx, endpoints.UpdateCurrentUser, options)).
		Operation("updateCurrentUser").
		Doc("Update current login user information (full).").
		Reads(endpoint.UpdateUserRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.PATCH("").
		To(NewKitHTTPServer[endpoint.PatchCurrentUserRequest](ctx, endpoints.PatchCurrentUser, options)).
		Operation("patchCurrentUser").
		Doc("Update current login user information (increment).").
		Reads(endpoint.PatchCurrentUserRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.GET("totp/secret").
		To(NewKitHTTPServer[endpoint.CreateTOTPSecretRequest](ctx, endpoints.CreateTOTPSecret, options)).
		Operation("getTOTPSecret").
		Doc("get TOTP Secret").
		Params(StructToQueryParams(endpoint.CreateTOTPSecretRequest{})...).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(global.MetaSensitiveData, true).
		Metadata(global.MetaNeedLogin, false).
		Returns(200, "OK", endpoint.CreateTOTPSecretResponse{}),
	)

	v1ws.Route(v1ws.POST("totp").
		To(NewKitHTTPServer[endpoint.CreateTOTPRequest](ctx, endpoints.CreateTOTP, options)).
		Operation("bindingTOTP").
		Doc("binding TOTP Secret").
		Reads(endpoint.CreateTOTPRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.POST("/activateAccount").
		To(NewKitHTTPServer[endpoint.ActivateAccountRequest](ctx, endpoints.ActivateAccount, options)).
		Operation("activateAccount").
		Reads(endpoint.ActivateAccountRequest{}).
		Doc("Activate the user.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(global.MetaSensitiveData, true).
		Metadata(global.MetaNeedLogin, false).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.POST("/forgotPassword").
		To(NewKitHTTPServer[endpoint.ForgotUserPasswordRequest](ctx, endpoints.ForgotPassword, options)).
		Operation("forgotPassword").
		Doc("Forgot the user password.").
		Reads(endpoint.ForgotUserPasswordRequest{}).
		Metadata(global.MetaNeedLogin, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.POST("/resetPassword").
		To(NewKitHTTPServer[endpoint.ResetUserPasswordRequest](ctx, endpoints.ResetPassword, options)).
		Operation("resetPassword").
		Reads(endpoint.ResetUserPasswordRequest{}).
		Doc("Reset the user password.").
		Metadata(global.MetaNeedLogin, false).
		Metadata(global.MetaSensitiveData, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.POST("/sendLoginCaptcha").
		To(NewKitHTTPServer[endpoint.SendLoginCaptchaRequest](ctx, endpoints.SendLoginCaptcha, options)).
		Operation("sendLoginCaptcha").
		Reads(endpoint.SendLoginCaptchaRequest{}).
		Doc("Send login code.").
		Metadata(global.MetaNeedLogin, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.SendLoginCaptchaResponse{}),
	)

	v1ws.Route(v1ws.GET("/sessions").
		To(NewKitHTTPServer[endpoint.GetSessionsRequest](ctx, endpoints.GetCurrentUserSessions, options)).
		Operation("getCurrentUserSessions").
		Doc("Get current user session list.").
		Params(StructToQueryParams(endpoint.GetSessionsRequest{})...).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetSessionsResponse{}),
	)

	v1ws.Route(v1ws.DELETE("/sessions/{id}").
		To(NewKitHTTPServer[endpoint.DeleteSessionRequest](ctx, endpoints.DeleteCurrentUserSession, options)).
		Operation("deleteCurrentUserSession").
		Doc("Delete current user a session.").
		Param(v1ws.PathParameter("id", "identifier of the session").DataType("string").Required(true)).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	return tag, []*restful.WebService{v1ws}
}

func PermissionService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "permissions", Description: "permissions service"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.GET("").
		To(NewKitHTTPServer[endpoint.GetPermissionsRequest](ctx, endpoints.GetPermissions, options)).
		Operation("getPermissions").
		Doc("Get permission list.").
		Params(StructToQueryParams(endpoint.GetPermissionsRequest{})...).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetPermissionsResponse{}),
	)
	return tag, []*restful.WebService{v1ws}
}

func RoleService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "roles", Description: "role service"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.GET("").
		To(NewKitHTTPServer[endpoint.GetRolesRequest](ctx, endpoints.GetRoles, options)).
		Operation("getRoles").
		Doc("Get role list.").
		Params(StructToQueryParams(endpoint.GetRolesRequest{})...).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetRolesResponse{}),
	)
	v1ws.Route(v1ws.DELETE("").
		To(NewKitHTTPServer[endpoint.DeleteRolesRequest](ctx, endpoints.DeleteRoles, options)).
		Operation("deleteRoles").
		Doc("Batch delete roles.").
		Reads(endpoint.DeleteRolesRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseTotalResponse{}),
	)
	v1ws.Route(v1ws.POST("").
		To(NewKitHTTPServer[endpoint.CreateRoleRequest](ctx, endpoints.CreateRole, options)).
		Operation("createRole").
		Doc("Create role.").
		Reads(endpoint.CreateRoleRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)
	v1ws.Route(v1ws.PUT("/{id}").
		To(NewKitHTTPServer[endpoint.UpdateRoleRequest](ctx, endpoints.UpdateRole, options)).
		Operation("updateRole").
		Doc("Update role information (full).").
		Param(v1ws.PathParameter("id", "identifier of the role").DataType("string")).
		Reads(endpoint.UpdateRoleRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	v1ws.Route(v1ws.DELETE("/{id}").
		To(NewKitHTTPServer[endpoint.DeleteRoleRequest](ctx, endpoints.DeleteRole, options)).
		Operation("deleteRole").
		Doc("删除角色").
		Param(v1ws.PathParameter("id", "identifier of the role").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	return tag, []*restful.WebService{v1ws}
}

func ConfigService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "config", Description: "config service"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.GET("security").
		To(NewKitHTTPServer[struct{}](ctx, endpoints.GetSecurityConfig, options)).
		Operation("getSecurityConfig").
		Doc("Obtain Security Configuration.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetSecurityConfigResponse{}),
	)
	v1ws.Route(v1ws.PATCH("security").
		To(NewKitHTTPServer[endpoint.PatchSecurityConfigRequest](ctx, endpoints.PatchSecurityConfig, options)).
		Operation("patchSecurityConfig").
		Doc("Update Security Configuration (Incremental).").
		Reads(endpoint.PatchSecurityConfigRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.BaseResponse{}),
	)

	return tag, []*restful.WebService{v1ws}
}

func EventService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "events", Description: "event service"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.GET("").
		To(NewKitHTTPServer[endpoint.GetEventsRequest](ctx, endpoints.GetEvents, options)).
		Operation("getEvents").
		Params(StructToQueryParams(endpoint.GetEventsRequest{})...).
		Doc("Get events.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetEventsResponse{}),
	)

	v1ws.Route(v1ws.GET("logs").
		To(NewKitHTTPServer[endpoint.GetEventLogsRequest](ctx, endpoints.GetEventLogs, options)).
		Operation("getEventLogs").
		Params(StructToQueryParams(endpoint.GetEventLogsRequest{})...).
		Doc("Get event logs.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", endpoint.GetEventLogsResponse{}),
	)

	return tag, []*restful.WebService{v1ws}
}

func GlobalService(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	tag := spec.Tag{TagProps: spec.TagProps{Name: "global", Description: "Global service"}}
	tags := []string{tag.Name}
	v1ws := NewWebService(RootPath, schema.GroupVersion{Group: tag.Name, Version: "v1"}, tag.Description)
	v1ws.Filter(HTTPAuthenticationFilter(endpoints))

	v1ws.Route(v1ws.GET("config").
		To(NewKitHTTPServer[struct{}](ctx, endpoints.GetGlobalConfig, options)).
		Operation("getGlobalConfig").
		Doc("Get global config.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(global.MetaNeedLogin, false).
		Returns(200, "OK", endpoint.GlobalConfigResponse{}),
	)

	return tag, []*restful.WebService{v1ws}
}

func XXLJobService(gv schema.GroupVersion) func(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
	return func(ctx context.Context, options []httptransport.ServerOption, endpoints endpoint.Set) (spec.Tag, []*restful.WebService) {
		tag := spec.Tag{TagProps: spec.TagProps{Name: gv.Group, Description: "xxl Job service."}}
		tags := []string{tag.Name}
		v1ws := NewWebService(RootPath, gv, tag.Description)

		v1ws.Route(v1ws.POST("/run").
			To(NewSimpleKitHTTPServer[endpoint.RunJobRequest](ctx, endpoints.ExecutorRunJob, DecodeHTTPRequest[endpoint.RunJobRequest], endpoint.DecodeJobHTTPResponse, options)).
			Operation("runJob").
			Consumes(restful.MIME_JSON).
			Reads(endpoint.RunJobRequest{}).
			Doc("run a job from xxljob").
			Metadata(restfulspec.KeyOpenAPITags, tags).
			Returns(200, "OK", endpoint.BaseJobResponse{}),
		)

		v1ws.Route(v1ws.POST("/beat").
			To(NewSimpleKitHTTPServer[endpoint.RunJobRequest](ctx, endpoints.ExecutorBeat, DecodeHTTPRequest[any], endpoint.DecodeJobHTTPResponse, options)).
			Operation("jobBeat").
			Consumes(restful.MIME_JSON).
			Doc("check executor is beat").
			Metadata(restfulspec.KeyOpenAPITags, tags).
			Returns(200, "OK", endpoint.BaseJobResponse{}),
		)
		v1ws.Route(v1ws.POST("/idleBeat").
			To(NewSimpleKitHTTPServer[endpoint.JobIdleBeatRequest](ctx, endpoints.ExecutorIdleBeat, DecodeHTTPRequest[endpoint.JobIdleBeatRequest], endpoint.DecodeJobHTTPResponse, options)).
			Operation("jobIdle").
			Consumes(restful.MIME_JSON).
			Doc("check executor is idle").
			Metadata(restfulspec.KeyOpenAPITags, tags).
			Returns(200, "OK", endpoint.BaseJobResponse{}),
		)

		v1ws.Route(v1ws.POST("/kill").
			To(NewSimpleKitHTTPServer[endpoint.KillJobRequest](ctx, endpoints.ExecutorKill, DecodeHTTPRequest[endpoint.KillJobRequest], endpoint.DecodeJobHTTPResponse, options)).
			Operation("runJob").
			Consumes(restful.MIME_JSON).
			Reads(endpoint.KillJobRequest{}).
			Doc("kill a job").
			Metadata(restfulspec.KeyOpenAPITags, tags).
			Returns(200, "OK", endpoint.BaseJobResponse{}),
		)
		v1ws.Route(v1ws.POST("/log").
			To(NewSimpleKitHTTPServer[endpoint.JobLogRequest](ctx, endpoints.ExecutorLog, DecodeHTTPRequest[endpoint.JobLogRequest], endpoint.DecodeJobHTTPResponse, options)).
			Operation("getJobLog").
			Consumes(restful.MIME_JSON).
			Reads(endpoint.JobLogRequest{}).
			Doc("get a job log").
			Metadata(restfulspec.KeyOpenAPITags, tags).
			Returns(200, "OK", endpoint.JobLogResponse{}),
		)
		return tag, []*restful.WebService{v1ws}
	}
}
