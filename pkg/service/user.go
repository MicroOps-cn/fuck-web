package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	logs "github.com/MicroOps-cn/fuck/log"
	w "github.com/MicroOps-cn/fuck/wrapper"
	"github.com/go-kit/log/level"
	uuid "github.com/satori/go.uuid"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck-web/pkg/service/opts"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/sign"
)

// GetUsers
//
//	@Description[en-US]: Get user list.
//	@Description[zh-CN]: 获取用户列表。
//	@param ctx       context.Context
//	@param keywords  string
//	@param status    models.UserMeta_UserStatus
//	@param current   int64
//	@param pageSize  int64
//	@return total    int64
//	@return users    []*models.User
//	@return err      error
func (s Set) GetUsers(ctx context.Context, keywords string, status models.UserMeta_UserStatus, current, pageSize int64) (total int64, users models.Users, err error) {
	total, users, err = s.GetUserService().GetUsers(ctx, keywords, status, current, pageSize)
	if err != nil {
		return total, users, err
	}
	if exts, err := s.commonService.GetUsersExtendedData(ctx, users.Id()); err == nil {
		for _, ext := range exts {
			users.GetById(ext.UserId).LoginTime = &ext.LoginTime
		}
	}
	if roles, err := s.commonService.GetUsersRole(ctx, users.Id()); err == nil {
		for _, role := range roles {
			fmt.Println(role, users.GetById(role.UserId))
			users.GetById(role.UserId).RoleId = role.RoleId
			users.GetById(role.UserId).Role = role.Role
		}
	}
	return total, users, nil
}

// PatchUsers
//
//	@Description[en-US]: Incrementally update information of multiple users.
//	@Description[zh-CN]: 增量更新多个用户的信息。
//	@param ctx 		context.Context
//	@param patch 	[]map[string]interface{}
//	@return count	int64
//	@return err		error
func (s Set) PatchUsers(ctx context.Context, patch []map[string]interface{}) (total int64, err error) {
	return s.GetUserService().PatchUsers(ctx, patch)
}

// DeleteUsers
//
//	@Description[en-US]: Delete users in batch.
//	@Description[zh-CN]: 批量删除用户。
//	@param ctx 		context.Context
//	@param ids 		[]string
//	@return count	int64
//	@return err		error
func (s Set) DeleteUsers(ctx context.Context, id []string) (total int64, err error) {
	return s.GetUserService().DeleteUsers(ctx, id)
}

// UpdateUser
//
//	@Description[en-US]: Update user information.
//	@Description[zh-CN]: 更新用户信息.
//	@param ctx	context.Context
//	@param user	*models.User
//	@param updateColumns	...string
//	@return err	error
func (s Set) UpdateUser(ctx context.Context, user *models.User, updateColumns ...string) (err error) {
	if err = s.GetUserService().UpdateUser(ctx, user, updateColumns...); err != nil {
		return err
	}
	if len(updateColumns) == 0 || w.Has(updateColumns, "role_id", w.StringEqual) {
		return s.commonService.PatchUserRole(ctx, user.Id, user.Role, user.RoleId)
	}
	return
}

// GetUserInfo
//
//	@Description[en-US]: Obtain user information through ID or username.
//	@Description[zh-CN]: 通过ID或用户名获取用户信息。
//	@param ctx 	context.Context
//	@param id 	string
//	@param username 	string
//	@return userDetail	*models.User
//	@return err	error
func (s Set) GetUserInfo(ctx context.Context, id string, username string) (user *models.User, err error) {
	return s.GetUserService().GetUserInfo(ctx, id, username)
}

// GetUser
//
//	@Description[en-US]: Get user info.
//	@Description[zh-CN]: 获取用户信息
//	@param ctx 	context.Context
//	@param options 	opts.WithGetUserOptions
//	@return userDetail	*models.User
//	@return err	error
func (s Set) GetUser(ctx context.Context, options ...opts.WithGetUserOptions) (user *models.User, err error) {
	o := opts.NewGetUserOptions(options...)
	if o.Err != nil {
		return nil, err
	}
	user, err = s.userService.GetUser(ctx, o)
	if err != nil {
		return nil, err
	}
	if role, err := s.commonService.GetUserRole(ctx, user.Id); err == nil {
		user.Role, user.RoleId = role.Name, role.Id
	} else {
		return nil, err
	}
	if o.Ext {
		user.ExtendedData, err = s.commonService.GetUserExtendedData(ctx, user.Id)
		if user.ExtendedData == nil {
			user.ExtendedData = new(models.UserExt)
		}
	}
	return user, err
}

// CreateUser
//
//	@Description[en-US]: Create a user.
//	@Description[zh-CN]: 创建用户。
//	@param ctx 	context.Context
//	@param user 	*models.User
//	@return err	error
func (s Set) CreateUser(ctx context.Context, user *models.User) (err error) {
	if len(user.Username) == 0 {
		return errors.ParameterError("username is null")
	}
	if err = s.GetUserService().CreateUser(ctx, user); err != nil {
		return err
	}
	return s.commonService.PatchUserRole(ctx, user.Id, "", user.RoleId)
}

// PatchUser
//
//	@Description[en-US]: Incremental update user.
//	@Description[zh-CN]: 增量更新用户。
//	@param ctx 	context.Context
//	@param user 	map[string]interface{}
//	@return err	error
func (s Set) PatchUser(ctx context.Context, user map[string]interface{}) (err error) {
	id, ok := user["id"].(string)
	if !ok {
		return errors.ParameterError("unknown id")
	}
	if err = s.GetUserService().PatchUser(ctx, user); err != nil {
		return err
	}

	if roleId, ok := user["role_id"].(string); ok {
		return s.commonService.PatchUserRole(ctx, id, "", roleId)
	} else if roleName, ok := user["role"].(string); ok {
		return s.commonService.PatchUserRole(ctx, id, roleName, "")
	}
	return nil
}

// PatchUserExtData
//
//	@Description[en-US]: Incremental update user.
//	@Description[zh-CN]: 增量更新用户扩展信息。
//	@param ctx 	context.Context
//	@param id 	string
//	@param patch 	map[string]interface{}
//	@return err	error
func (s Set) PatchUserExtData(ctx context.Context, userId string, patch map[string]interface{}) (err error) {
	return s.commonService.PatchUserExtData(ctx, userId, patch)
}

// DeleteUser
//
//	@Description[en-US]: Delete a user.
//	@Description[zh-CN]: 删除用户。
//	@param ctx 	context.Context
//	@param id 	string
//	@return error
func (s Set) DeleteUser(ctx context.Context, id string) (err error) {
	return s.GetUserService().DeleteUser(ctx, id)
}

// VerifyPasswordById
//
//	@Description[en-US]: Verify the user's password through ID.
//	@Description[zh-CN]: 通过ID验证用户密码。
//	@param ctx 	context.Context
//	@param id 	string
//	@param password 	string
//	@return users	[]*models.User
func (s Set) VerifyPasswordById(ctx context.Context, userId, password string, allowPasswordExpired bool) (user *models.User) {
	begin := time.Now().UTC()
	var err error
	defer func() {
		var username string
		if user != nil {
			username = user.Username
		} else if err == nil {
			err = fmt.Errorf("Failed to verify password. ")
		}
		eventId, message, status, took := GetEventMeta(ctx, "VerifyPassword", begin, err)
		if e := s.PostEventLog(ctx, eventId, userId, username, "", "VerifyPassword", message, status, took); e != nil {
			level.Error(logs.GetContextLogger(ctx)).Log("failed to post event log", "err", e)
		}
	}()
	user = s.GetUserService().VerifyPasswordById(ctx, userId, password)
	if user == nil {
		return nil
	}
	return user
}

// VerifyPassword
//
//	@Description[en-US]: Verify password for user.
//	@Description[zh-CN]: 验证用户密码。
//	@param ctx 	context.Context
//	@param username 	string
//	@param password 	string
//	@return users	[]*models.User
func (s Set) VerifyPassword(ctx context.Context, username string, password string, allowPasswordExpired bool) (user *models.User, err error) {
	begin := time.Now().UTC()
	defer func() {
		var userId string
		if user != nil && len(user.Id) > 0 {
			userId = user.Id
		} else if user != nil && err == nil {
			err = fmt.Errorf("Failed to verify password. ")
		}
		eventId, message, status, took := GetEventMeta(ctx, "VerifyPassword", begin, err)
		if e := s.PostEventLog(ctx, eventId, userId, username, "", "VerifyPassword", message, status, took); e != nil {
			level.Error(logs.GetContextLogger(ctx)).Log("failed to post event log", "err", e)
		}
	}()
	var ts int64
	var counterSeed string
	failedSec, failedThreshold := config.GetRuntimeConfig().GetPasswordFailedLockConfig()

	if failedSec > 0 && failedThreshold > 0 {
		nowTs := time.Now().Unix()
		ts = nowTs - nowTs%failedSec
		counterSeed = fmt.Sprintf("LOGIN:%s:%d", username, ts)
		count, err := s.sessionService.GetCounter(ctx, counterSeed)
		if err != nil {
			level.Error(logs.GetContextLogger(ctx)).Log("msg", "Failed to obtain password counter.", "err", err)
			return nil, errors.NewServerError(http.StatusInternalServerError, "System error: Please contact the administrator.", errors.CodeSystemError)
		}
		if count >= failedThreshold {
			return nil, errors.NewServerError(http.StatusOK, "The number of password errors has reached the threshold. ", errors.CodeTooManyLoginFailures)
		}
	}
	user = s.userService.VerifyPassword(ctx, username, password)
	if user == nil {
		if ts > 0 && len(counterSeed) > 0 {
			expir := time.Unix(ts+failedSec, 0)
			if err = s.sessionService.Counter(ctx, counterSeed, &expir); err != nil {
				level.Error(logs.GetContextLogger(ctx)).Log("msg", "Failed to write password failure counter.")
			}
		}
		return nil, errors.NewServerError(http.StatusOK, "Wrong user name or password. ", errors.CodeInvalidCredentials)
	}
	if user.ExtendedData == nil {
		user.ExtendedData, err = s.commonService.GetUserExtendedData(ctx, user.Id)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (s Set) VerifyUserStatus(ctx context.Context, user *models.User, allowPasswordExpired bool) (err error) {
	logger := logs.GetContextLogger(ctx)
	switch user.Status {
	case models.UserMeta_normal:
	case models.UserMeta_user_inactive:
		return errors.NewServerError(http.StatusOK, "The user is disabled due to inactivity. Please contact administrator.", errors.CodeUserInactive)
	case models.UserMeta_password_expired:
		if !allowPasswordExpired {
			return errors.NewServerError(http.StatusOK, "Your password has expired. Please change your password and log in again.", errors.CodeUserNeedResetPassword)
		}
	case models.UserMeta_disabled:
		return errors.NewServerError(http.StatusOK, "The user status is abnormal. Please contact the administrator.", errors.CodeUserDisable)
	default:
		return errors.NewServerError(http.StatusOK, "Unknown user status.", errors.CodeUserStatusUnknown)
	}
	if user.ExtendedData == nil || len(user.ExtendedData.UserId) == 0 {
		user.ExtendedData, err = s.commonService.GetUserExtendedData(ctx, user.Id)
		if err != nil {
			return errors.WithServerError(http.StatusInternalServerError, err, "Failed to obtain user. ")
		}
	}

	if accountInactiveLock := time.Duration(config.GetRuntimeConfig().Security.AccountInactiveLock) * time.Hour * 24; accountInactiveLock > 0 {
		if time.Since(user.ExtendedData.LoginTime) > accountInactiveLock &&
			time.Since(user.ExtendedData.PasswordModifyTime) > accountInactiveLock &&
			time.Since(user.ExtendedData.ActivationTime) > accountInactiveLock {
			user.Status = models.UserMeta_user_inactive
			if err = s.UpdateUser(ctx, user, "status"); err != nil {
				level.Error(logger).Log("msg", "failed to update user status", "err", err)
			}
			return errors.NewServerError(http.StatusOK, "The user is disabled due to inactivity. Please contact administrator.", errors.CodeUserInactive)
		}
	}
	if passwordExpireTime := config.GetRuntimeConfig().Security.PasswordExpireTime; passwordExpireTime > 0 {
		if time.Since(user.ExtendedData.PasswordModifyTime) > time.Duration(passwordExpireTime)*time.Hour*24 && !allowPasswordExpired {
			user.Status = models.UserMeta_password_expired
			if err := s.UpdateUser(ctx, user, "status"); err != nil {
				level.Error(logger).Log("msg", "failed to update user status", "err", err)
			}
			return errors.NewServerError(http.StatusOK, "Your password has expired. Please change the password and log in again.", errors.CodeUserNeedResetPassword)
		}
	}
	if err = s.commonService.UpdateLoginTime(ctx, user.Id); err != nil {
		level.Error(logger).Log("msg", "failed to update user login time", "err", err)
	}
	return nil
}

// Authentication
//
//	@Description[en-US]: Authenticate the user's request.
//	@Description[zh-CN]: 对用户请求进行身份认证。
//	@param ctx 	context.Context
//	@param method 	models.AuthMeta_Method
//	@param algorithm 	sign.AuthAlgorithm
//	@param key 	string
//	@param secret 	string
//	@param payload 	string
//	@param signStr 	string
//	@return user	[]*models.User
//	@return err	error
func (s Set) Authentication(ctx context.Context, method models.AuthMeta_Method, algorithm sign.AuthAlgorithm, key, secret, payload, signStr string) (user *models.User, err error) {
	if method == models.AuthMeta_basic {
		if _, err = uuid.FromString(key); err != nil {
			if config.Get().GetSecurity().DisableLoginForm {
				return nil, errors.ParameterError("unsupported login type")
			}
			user, err = s.VerifyPassword(ctx, key, secret, false)
			if err != nil {
				return nil, err
			} else if err = s.VerifyUserStatus(ctx, user, false); err != nil {
				return nil, err
			}
			return user, err
		}
	}
	return nil, errors.ParameterError("unknown auth request")
}

func (s Set) GetUserInfoByUsernameAndEmail(ctx context.Context, username, email string) (users *models.User, err error) {
	if len(username) == 0 {
		return nil, errors.LackParameterError("username")
	}
	if len(email) == 0 {
		return nil, errors.LackParameterError("email")
	}
	return s.GetUserService().GetUserInfoByUsernameAndEmail(ctx, username, email)
}

func (s Set) CreateTOTP(ctx context.Context, id string, secret string) error {
	return s.commonService.CreateTOTP(ctx, id, secret)
}

func (s Set) GetTOTPSecrets(ctx context.Context, ids []string) ([]string, error) {
	return s.commonService.GetTOTPSecrets(ctx, ids)
}

func (s Set) UpdateUserSession(ctx context.Context, userId string) (err error) {
	newUser, err := s.GetUser(ctx, opts.WithUserId(userId), opts.WithUserExt)
	if err != nil {
		return err
	}
	logger := logs.GetContextLogger(ctx)
	if err != nil {
		level.Warn(logger).Log("msg", "Failed to update current user cache info: can't get user info.", "err", err)
	} else {
		role, err := s.GetUserRole(ctx, userId)
		if err == nil {
			newUser.RoleId = role.Id
			newUser.Role = role.Name
		} else if !errors.IsNotFount(err) {
			level.Error(logger).Log("msg", "failed to get app role", "err", err)
		}
		var sessions []*models.Token
		var maxCount, count, current int64
		for count <= maxCount {
			current++
			maxCount, sessions, err = s.sessionService.GetSessions(ctx, userId, current, 100)
			if err != nil {
				return err
			} else if len(sessions) == 0 {
				return nil
			}
			count += int64(len(sessions))
			for _, tk := range sessions {
				var oldUser models.User
				if err := tk.To(&oldUser); err != nil {
					return err
				} else if oldUser.ExtendedData == nil {
					oldUser.ExtendedData = new(models.UserExt)
				}
				newUser.ExtendedData.LoginTime = oldUser.ExtendedData.LoginTime
				rawData, err := json.Marshal(newUser)
				if err != nil {
					return err
				}
				tk.Data = rawData
				if err = s.sessionService.UpdateToken(ctx, tk); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s Set) ResetPassword(ctx context.Context, id string, password string) (err error) {
	logger := logs.GetContextLogger(ctx)
	begin := time.Now().UTC()
	defer func() {
		eventId, message, status, took := GetEventMeta(ctx, "ResetPassword", begin, err)
		if e := s.PostEventLog(ctx, eventId, id, "", "", "ResetPassword", message, status, took); e != nil {
			level.Error(logs.GetContextLogger(ctx)).Log("failed to post event log", "err", e)
		}
	}()
	userExtendedData, err := s.commonService.GetUserExtendedData(ctx, id)
	if err != nil {
		return errors.WithServerError(http.StatusInternalServerError, err, "Failed to obtain user. ")
	}
	if accountInactiveLock := time.Duration(config.GetRuntimeConfig().Security.AccountInactiveLock) * time.Hour * 24; accountInactiveLock > 0 {
		if time.Since(userExtendedData.LoginTime) > accountInactiveLock &&
			time.Since(userExtendedData.PasswordModifyTime) > accountInactiveLock &&
			time.Since(userExtendedData.ActivationTime) > accountInactiveLock {
			return errors.NewServerError(http.StatusForbidden, "The user is disabled due to inactivity. Please contact administrator.", errors.CodeUserInactive)
		}
	}

	if err = s.commonService.VerifyWeakPassword(ctx, password); err != nil {
		return err
	} else if err = s.commonService.VerifyAndRecordHistoryPassword(ctx, id, password); err != nil {
		return err
	} else if err = s.userService.ResetPassword(ctx, id, password); err != nil {
		return fmt.Errorf("failed to reset password: %s", err)
	}
	if err = s.commonService.PatchUserExtData(ctx, id, map[string]interface{}{
		"password_modify_time": time.Now().UTC(),
	}); err != nil {
		level.Error(logger).Log("err", err, "msg", "failed to update `password_modify_time` and `login_time`")
		return fmt.Errorf("The password was successfully modified, but a slight error was encountered. ")
	}
	return nil
}

func (s Set) GetUserRole(ctx context.Context, id string) (*models.Role, error) {
	return s.commonService.GetUserRole(ctx, id)
}
