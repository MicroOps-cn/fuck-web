declare namespace API {
  type ActivateAccountRequest = {
    newPassword: string;
    token?: string;
    userId: string;
  };

  type BaseListResponse = {
    current: number;
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type BaseResponse = {
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    traceId: string;
  };

  type BaseTotalResponse = {
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    total: number;
    traceId: string;
  };

  type CreateRoleRequest = {
    description?: string;
    name: string;
    permission?: string[];
  };

  type CreateTOTPRequest = {
    firstCode: string;
    secondCode: string;
    token: string;
  };

  type CreateTOTPSecretResponse = {
    data?: CreateTOTPSecretResponseData;
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    traceId: string;
  };

  type CreateTOTPSecretResponseData = {
    secret: string;
    token: string;
  };

  type CreateUserRequest = {
    avatar?: string;
    email?: string;
    fullName?: string;
    isDelete?: boolean;
    phoneNumber?: string;
    roleId?: string;
    status?: UserMetaUserStatus;
    username: string;
  };

  type currentUserEventLogsParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
    eventId: string;
  };

  type currentUserEventsParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
    action?: string;
    startTime: string;
    endTime: string;
  };

  type deleteCurrentUserSessionParams = {
    /** identifier of the session */
    id: string;
  };

  type deleteRoleParams = {
    /** identifier of the role */
    id: string;
  };

  type DeleteRoleRequest = {
    id: string;
  };

  type deleteSessionParams = {
    /** identifier of the session */
    id: string;
  };

  type deleteUserParams = {
    /** identifier of the user */
    id: string;
  };

  type DeleteUserRequest = {
    id: string;
  };

  type downloadFileParams = {
    /** identifier of the file */
    id: string;
  };

  type Event = {
    action: string;
    client_ip: string;
    createTime: string;
    id: string;
    location: string;
    message: string;
    status: string;
    took: number;
    updateTime: string;
    userId: string;
    username: string;
  };

  type EventLog = {
    createTime: string;
    id: string;
    log: string;
    updateTime: string;
    userId: string;
  };

  type FileUploadResponse = {
    data: Record<string, any>;
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    total: number;
    traceId: string;
  };

  type ForgotUserPasswordRequest = {
    email: string;
    username: string;
  };

  type GetCurrentUserEventLogsResponse = {
    current: number;
    data?: EventLog[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type GetCurrentUserEventsResponse = {
    current: number;
    data?: Event[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type getCurrentUserSessionsParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
    userId?: string;
  };

  type getEventLogsParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
    eventId: string;
  };

  type GetEventLogsResponse = {
    current: number;
    data?: EventLog[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type getEventsParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
    username?: string;
    action?: string;
    startTime: string;
    endTime: string;
  };

  type GetEventsResponse = {
    current: number;
    data?: Event[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type getPermissionsParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
  };

  type GetPermissionsResponse = {
    current: number;
    data?: PermissionInfo[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type getRolesParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
  };

  type GetRolesResponse = {
    current: number;
    data?: RoleInfo[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type GetSecurityConfigResponse = {
    data?: RuntimeSecurityConfig;
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    traceId: string;
  };

  type getSessionsParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
    userId?: string;
  };

  type GetSessionsResponse = {
    current: number;
    data?: SessionInfo[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type getTOTPSecretParams = {
    token?: any;
  };

  type getUserInfoParams = {
    /** identifier of the user */
    id: string;
  };

  type GetUserResponse = {
    data?: UserInfo;
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    traceId: string;
  };

  type getUsersParams = {
    pageSize?: number;
    current?: number;
    keywords?: string;
    status?: UserMetaUserStatus;
  };

  type GetUsersResponse = {
    current: number;
    data?: UserInfo[];
    errorCode?: string;
    errorMessage?: string;
    pageSize: number;
    success: boolean;
    total: number;
    traceId: string;
  };

  type GlobalConfig = {
    admin_url?: string;
    copyright?: string;
    defaultLoginType: LoginType;
    external_url?: string;
    loginType: GlobalLoginType[];
    logo?: string;
    subTitle?: string;
    title?: string;
    version?: string;
  };

  type GlobalConfigResponse = {
    data?: GlobalConfig;
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    traceId: string;
  };

  type GlobalLoginType = {
    autoLogin?: boolean;
    autoRedirect?: boolean;
    icon?: string;
    id?: string;
    name?: string;
    type: LoginType;
  };

  type LoginType =
    | 'normal'
    | 0
    | 'mfa_totp'
    | 1
    | 'mfa_email'
    | 2
    | 'mfa_sms'
    | 3
    | 'email'
    | 4
    | 'sms'
    | 5
    | 'oauth2'
    | 6
    | 'enable_mfa_totp'
    | 10
    | 'enable_mfa_email'
    | 11
    | 'enable_mfa_sms'
    | 12;

  type PasswordComplexity = 'unsafe' | 0 | 'general' | 1 | 'safe' | 2 | 'very_safe' | 3;

  type PatchCurrentUserRequest = {
    email_as_mfa?: boolean;
    sms_as_mfa?: boolean;
    totp_as_mfa?: boolean;
  };

  type PatchSecurityConfigRequest = {
    accountInactiveLock?: number;
    forceEnableMfa?: boolean;
    loginSessionInactivityTime?: number;
    loginSessionMaxTime?: number;
    passwordComplexity?: PasswordComplexity;
    passwordExpireTime?: number;
    passwordFailedLockDuration?: number;
    passwordFailedLockThreshold?: number;
    passwordHistory?: number;
    passwordMinLength?: number;
  };

  type patchUserParams = {
    /** identifier of the user */
    id: string;
  };

  type PatchUserRequest = {
    id?: string;
    isDelete?: boolean;
    status?: UserMetaUserStatus;
  };

  type PatchUserResponse = {
    User: string;
  };

  type PermissionInfo = {
    createTime: string;
    description?: string;
    enableAuth?: boolean;
    id: string;
    name?: string;
    parentId?: string;
    updateTime: string;
  };

  type ResetUserPasswordRequest = {
    newPassword: string;
    oldPassword?: string;
    token?: string;
    userId: string;
    username?: string;
  };

  type RoleInfo = {
    createTime: string;
    describe?: string;
    id: string;
    name: string;
    permission?: PermissionInfo[];
    updateTime: string;
  };

  type RuntimeSecurityConfig = {
    accountInactiveLock: number;
    forceEnableMfa: boolean;
    loginSessionInactivityTime: number;
    loginSessionMaxTime: number;
    passwordComplexity: PasswordComplexity;
    passwordExpireTime: number;
    passwordFailedLockDuration: number;
    passwordFailedLockThreshold: number;
    passwordHistory: number;
    passwordMinLength: number;
  };

  type SendActivationMailRequest = {
    userId: string;
  };

  type SendLoginCaptchaRequest = {
    email?: string;
    phone?: string;
    type: LoginType;
    username?: string;
  };

  type SendLoginCaptchaResponse = {
    data?: SendLoginCaptchaResponseData;
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    traceId: string;
  };

  type SendLoginCaptchaResponseData = {
    token?: string;
  };

  type SessionInfo = {
    createTime: string;
    expiry: string;
    id: string;
    lastSeen?: string;
  };

  type updateRoleParams = {
    /** identifier of the role */
    id: string;
  };

  type UpdateRoleRequest = {
    description?: string;
    id: string;
    name: string;
    permission?: string[];
  };

  type updateUserParams = {
    /** identifier of the user */
    id: string;
  };

  type UpdateUserRequest = {
    avatar?: string;
    email?: string;
    fullName?: string;
    id: string;
    isDelete?: boolean;
    phoneNumber?: string;
    roleId?: string;
    status?: UserMetaUserStatus;
    username: string;
  };

  type uploadFileParams = {
    /** files */
    files?: string[];
  };

  type UserExt = {
    ForceMFA: boolean;
    activationTime: string;
    emailAsMFA: boolean;
    loginTime: string;
    passwordModifyTime: string;
    smsAsMFA: boolean;
    totpAsMFA: boolean;
    userId: string;
  };

  type UserInfo = {
    avatar?: string;
    createTime: string;
    email?: string;
    extendedData?: UserExt;
    fullName?: string;
    id: string;
    isDelete: boolean;
    loginTime?: string;
    phoneNumber?: string;
    role?: string;
    roleId?: string;
    status: UserMetaUserStatus;
    updateTime: string;
    username: string;
  };

  type UserLoginRequest = {
    autoLogin?: boolean;
    bindingToken?: string;
    code?: string;
    email?: string;
    firstCode?: string;
    password?: string;
    phone?: string;
    secondCode?: string;
    token?: string;
    type?: LoginType;
    username?: string;
  };

  type UserLoginResponse = {
    data?: UserLoginResponseData;
    errorCode?: string;
    errorMessage?: string;
    success: boolean;
    traceId: string;
  };

  type UserLoginResponseData = {
    email?: string;
    nextMethod: LoginType[];
    phone_number?: string;
    token?: string;
  };

  type UserMetaUserStatus =
    | 'normal'
    | 0
    | 'disabled'
    | 1
    | 'user_inactive'
    | 2
    | 'password_expired'
    | 4;

  type userOAuthLoginParams = {
    /** identifier of the oauth */
    id: string;
  };
}
