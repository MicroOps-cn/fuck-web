package gormservice

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	logs "github.com/MicroOps-cn/fuck/log"
	w "github.com/MicroOps-cn/fuck/wrapper"
	"github.com/go-kit/log/level"
	uuid "github.com/satori/go.uuid"
	gogorm "gorm.io/gorm"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck-web/pkg/service/opts"
	"github.com/MicroOps-cn/fuck-web/pkg/utils/sign"
)

// ResetPassword
//
//	@Description[en-US]: Reset User Password.
//	@Description[zh-CN]: 重置用户密码。
//	@param ctx       context.Context
//	@param id        string
//	@param password  string           : New password.
//	@return err      error
func (s UserService) ResetPassword(ctx context.Context, ids string, password string) error {
	conn := s.Session(ctx).Begin()
	defer conn.Callback()
	for _, id := range strings.Split(ids, ",") {
		u := models.User{Model: models.Model{Id: id}, Salt: w.M(uuid.NewV4()).Bytes(), Status: models.UserMeta_normal}
		u.Password = u.GenSecret(password)
		if err := conn.Select("password", "salt", "status").Where("status not in ?", []models.UserMeta_UserStatus{
			models.UserMeta_disabled,
		}).Updates(&u).Error; err != nil {
			return err
		}
	}

	return conn.Commit().Error
}

func (s UserService) Name() string {
	return s.name
}

const sqlGetUserAndRoleInfoById = `
SELECT 
    T4.id AS role_id, T4.name AS role, T1.*
FROM
    t_user T1
        LEFT JOIN
    t_app_user T2 ON T2.user_id = T1.id
        LEFT JOIN
    t_app T3 ON T3.id = T2.app_id 
        LEFT JOIN
    t_app_role T4 ON T2.role_id = T4.id
WHERE
    T1.id = ?
    AND T3.name = 'IDAS'
`

// VerifyPasswordById
//
//	@Description[en-US]: Verify the user's password through ID.
//	@Description[zh-CN]: 通过ID验证用户密码。
//	@param ctx 	context.Context
//	@param id 	string
//	@param password 	string
//	@return users	[]*models.User
func (s UserService) VerifyPasswordById(ctx context.Context, id, password string) *models.User {
	logger := logs.GetContextLogger(ctx)
	var user models.User
	if err := s.Session(ctx).Raw(sqlGetUserAndRoleInfoById, id).First(&user).Error; err != nil {
		if err == gogorm.ErrRecordNotFound {
			level.Debug(logger).Log("msg", "incorrect username", "id", id)
		} else {
			level.Error(logger).Log("msg", "unknown error", "id", id, "err", err)
		}
		return nil
	}
	if !bytes.Equal(user.GenSecret(password), user.Password) {
		level.Debug(logger).Log("msg", "incorrect password", "id", id)
		return nil
	}
	return &user
}

// VerifyPassword
//
//	@Description[en-US]: Verify password for user.
//	@Description[zh-CN]: 验证用户密码。
//	@param ctx 	context.Context
//	@param username 	string
//	@param password 	string
//	@return users	[]*models.User
func (s UserService) VerifyPassword(ctx context.Context, username string, password string) *models.User {
	logger := logs.GetContextLogger(ctx)
	var user models.User
	if err := s.Session(ctx).Where("(username = ? or email = ?) and delete_time is NULL", username, username).First(&user).Error; err != nil {
		if err == gogorm.ErrRecordNotFound {
			level.Debug(logger).Log("msg", "incorrect username", "username", username)
		} else {
			level.Error(logger).Log("msg", "unknown error", "username", username, "err", err)
		}
		return nil
	}
	if !bytes.Equal(user.GenSecret(password), user.Password) {
		level.Debug(logger).Log("msg", "incorrect password", "username", username)
		return nil
	}
	return &user
}

// GetUserInfoByUsernameAndEmail
//
//	@Description[en-US]: Use username or email to obtain user information.
//	@Description[zh-CN]: 使用用户名或email获取用户信息。
//	@param ctx           context.Context
//	@param username      string
//	@param email         string
//	@return user   *models.User
//	@return err          error
func (s UserService) GetUserInfoByUsernameAndEmail(ctx context.Context, username, email string) (user *models.User, err error) {
	user = new(models.User)
	query := s.Session(ctx).Where("username = ? and email = ? and delete_time is NULL", username, email)
	if err = query.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers
//
//	@Description[en-US]: Get user list.
//	@Description[zh-CN]: 获取用户列表。
//	@param ctx       context.Context
//	@param keywords  string
//	@param status    models.UserMeta_UserStatus
//	@param appId     string
//	@param current   int64
//	@param pageSize  int64
//	@return total    int64
//	@return users    []*models.User
//	@return err      error
func (s UserService) GetUsers(ctx context.Context, keywords string, status models.UserMeta_UserStatus, current, pageSize int64) (total int64, users []*models.User, err error) {
	conn := s.Session(ctx)
	query := conn.Model(&models.User{})
	if len(keywords) > 0 {
		keywords = fmt.Sprintf("%%%s%%", keywords)
		query = query.Where(
			conn.Where("username like ?", keywords).
				Or("email like ?", keywords).
				Or("phone_number like ?", keywords).
				Or("full_name like ?", keywords),
		)
	}
	if status != models.UserMetaStatusAll {
		query = query.Where("status", status)
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return 0, nil, err
	} else if err = query.Order("username,id").Limit(int(pageSize)).Offset(int((current - 1) * pageSize)).Find(&users).Error; err != nil {
		return 0, nil, err
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
func (s UserService) PatchUsers(ctx context.Context, patch []map[string]interface{}) (int64, error) {
	var patchCount int64
	tx := s.Session(ctx).Begin()
	defer tx.Rollback()
	updateQuery := tx.Model(&models.User{}).Select("is_delete", "status")
	var newPatch map[string]interface{}
	var newPatchIds []string
	for _, patchInfo := range patch {
		tmpPatch := map[string]interface{}{}
		var tmpPatchId string
		for name, value := range patchInfo {
			if name != "id" {
				tmpPatch[name] = value
			} else {
				tmpPatchId, _ = value.(string)
			}
		}
		if tmpPatchId == "" {
			return 0, errors.ParameterError("invalid id")
		} else if len(tmpPatch) == 0 {
			return 0, errors.ParameterError("update content is empty")
		}
		if len(newPatchIds) == 0 {
			newPatchIds = append(newPatchIds, tmpPatchId)
			newPatch = tmpPatch
		} else if reflect.DeepEqual(tmpPatch, newPatch) {
			newPatchIds = append(newPatchIds, tmpPatchId)
		} else {
			patched := updateQuery.Where("id in ?", newPatchIds).Updates(newPatch)
			if err := patched.Error; err != nil {
				return 0, err
			}
			patchCount = patched.RowsAffected
			newPatchIds = []string{}
			newPatch = map[string]interface{}{}
		}
	}
	if len(newPatchIds) > 0 {
		patched := updateQuery.Where("id in ?", newPatchIds).Updates(newPatch)
		if err := patched.Error; err != nil {
			return 0, err
		}
		patchCount = patched.RowsAffected
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return patchCount, nil
}

// DeleteUsers
//
//	@Description[en-US]: Delete users in batch.
//	@Description[zh-CN]: 批量删除用户。
//	@param ctx 		context.Context
//	@param ids 		[]string
//	@return count	int64
//	@return err		error
func (s UserService) DeleteUsers(ctx context.Context, id []string) (int64, error) {
	deleted := s.Session(ctx).Model(&models.User{}).Where("id in ?", id).Update("delete_time", time.Now().UTC())
	if err := deleted.Error; err != nil {
		return deleted.RowsAffected, err
	}
	return deleted.RowsAffected, nil
}

// UpdateUser
//
//	@Description[en-US]: Update user information.
//	@Description[zh-CN]: 更新用户信息.
//	@param ctx	context.Context
//	@param user	*models.User
//	@param updateColumns	...string
//	@return err	error
func (s UserService) UpdateUser(ctx context.Context, user *models.User, updateColumns ...string) (err error) {
	tx := s.Session(ctx).Begin()
	defer tx.Rollback()
	q := tx.Omit("create_time")
	if len(updateColumns) != 0 {
		q = q.Select(updateColumns)
	} else {
		q = q.Select("email", "phone_number", "full_name", "avatar", "status")
	}

	if err = q.Updates(&user).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// GetUser
//
//	@Description[en-US]: Get user info.
//	@Description[zh-CN]: 获取用户信息
//	@param ctx 	context.Context
//	@param options 	opts.GetUserOptions
//	@return userDetail	*models.User
//	@return err	error
func (s UserService) GetUser(ctx context.Context, o *opts.GetUserOptions) (*models.User, error) {
	conn := s.Session(ctx)
	var user models.User
	query := conn.Model(&models.User{})
	if len(o.Id) != 0 {
		query.Where("id = ?", o.Id)
	} else if len(o.Username) != 0 {
		query.Where("username = ?", o.Username)
	} else if len(o.Email) != 0 {
		query.Where("email = ?", o.Email)
	} else if len(o.PhoneNumber) > 0 {
		query.Where("phone_number = ?", o.PhoneNumber)
	}
	if err := query.First(&user).Error; err != nil {
		if err == gogorm.ErrRecordNotFound {
			return nil, errors.StatusNotFound("user")
		}
		return nil, err
	}
	return &user, nil
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
func (s UserService) GetUserInfo(ctx context.Context, id string, username string) (*models.User, error) {
	conn := s.Session(ctx)
	var user models.User
	query := conn.Model(&models.User{})
	if len(id) != 0 && len(username) != 0 {
		subQuery := query.Where("id = ?", id).Or("username = ?", username)
		query = query.Where(subQuery)
	} else if len(id) != 0 {
		query = query.Where("id = ?", id)
	} else if len(username) != 0 {
		query = query.Where("username = ?", username)
	} else {
		return nil, errors.ParameterError("require id or username")
	}
	if err := query.First(&user).Error; err != nil {
		if err == gogorm.ErrRecordNotFound {
			return nil, errors.StatusNotFound("user")
		}
		return nil, err
	}
	return &user, nil
}

func (s UserService) GetUsersById(ctx context.Context, id []string) (users models.Users, err error) {
	conn := s.Session(ctx)
	query := conn.Model(&models.User{}).Where("id in ?", id)
	if err = query.Find(&users).Error; err != nil {
		if err == gogorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return
}

// CreateUser
//
//	@Description[en-US]: Create a user.
//	@Description[zh-CN]: 创建用户。
//	@param ctx 	context.Context
//	@param user 	*models.User
//	@return err	error
func (s UserService) CreateUser(ctx context.Context, user *models.User) (err error) {
	conn := s.Session(ctx)
	if len(user.Password) != 0 {
		user.Salt = w.M(uuid.NewV4()).Bytes()
		user.Password = user.GenSecret()
	}
	return conn.Omit("role", "role_id").Create(user).Error
}

// PatchUser
//
//	@Description[en-US]: Incremental update user.
//	@Description[zh-CN]: 增量更新用户。
//	@param ctx 	context.Context
//	@param user 	map[string]interface{}
//	@return err	error
func (s UserService) PatchUser(ctx context.Context, patch map[string]interface{}) (err error) {
	if id, ok := patch["id"].(string); ok {
		tx := s.Session(ctx).Begin()
		delete(patch, "username")
		if err = tx.Model(&models.User{}).Where("id = ?", id).Updates(patch).Error; err != nil {
			return err
		}
		return tx.Commit().Error
	}
	return errors.ParameterError("id is null")
}

// DeleteUser
//
//	@Description[en-US]: Delete a user.
//	@Description[zh-CN]: 删除用户。
//	@param ctx 	context.Context
//	@param id 	string
//	@return error
func (s UserService) DeleteUser(ctx context.Context, id string) (err error) {
	_, err = s.DeleteUsers(ctx, []string{id})
	return err
}

func (c *CommonService) GetUserExtendedData(ctx context.Context, id string) (*models.UserExt, error) {
	conn := c.Session(ctx)
	var ext models.UserExt
	err := conn.Where("user_id = ?", id).First(&ext).Error
	if err == gogorm.ErrRecordNotFound {
		return nil, nil
	}
	if ext.EmailAsMFA || ext.SmsAsMFA || ext.TOTPAsMFA {
		ext.ForceMFA = true
	}
	return &ext, err
}

func (c *CommonService) GetUsersExtendedData(ctx context.Context, id []string) ([]*models.UserExt, error) {
	conn := c.Session(ctx)
	if len(id) == 0 {
		return nil, nil
	}
	var exts []*models.UserExt
	err := conn.Where("user_id in ?", id).Find(&exts).Error
	if err == gogorm.ErrRecordNotFound {
		return nil, nil
	}
	for _, ext := range exts {
		if ext.EmailAsMFA || ext.SmsAsMFA || ext.TOTPAsMFA {
			ext.ForceMFA = true
		}
	}
	return exts, err
}

func (c *CommonService) PatchUserExtData(ctx context.Context, id string, patch map[string]interface{}) error {
	conn := c.Session(ctx)
	ext := models.UserExt{UserId: id}
	if created := c.Session(ctx).FirstOrCreate(&ext); created.Error != nil {
		return created.Error
	}
	return conn.Model(&models.UserExt{}).Where("user_id = ?", id).Updates(patch).Error
}

func (c *CommonService) VerifyAndRecordHistoryPassword(ctx context.Context, id string, password string) error {
	passHisCount := int(config.GetRuntimeConfig().Security.PasswordHistory)
	if passHisCount == 0 {
		return nil
	}
	logger := logs.GetContextLogger(ctx)
	his := &models.UserPasswordHistory{UserId: id}
	if err := his.SetPassword(password); err != nil {
		return err
	}
	tx := c.Session(ctx).Begin()
	defer tx.Rollback()
	var lastHis []models.UserPasswordHistory
	if err := tx.Model(&models.UserPasswordHistory{}).Order("create_time DESC").Limit(passHisCount).Find(&lastHis).Error; err != nil {
		return fmt.Errorf("failed to obtain the last %d passwords. ", passHisCount)
	}
	for _, lh := range lastHis {
		if lh.Hash == his.Hash {
			return errors.NewServerError(400, fmt.Sprintf("Coincident with the last %d passwords used", passHisCount), errors.CodePasswordRepetition)
		}
	}
	if err := tx.Create(his).Error; err != nil {
		level.Error(logger).Log("msg", "Failed to create password history", "err", err)
	} else if len(lastHis) > 0 {
		if err = tx.Delete(&lastHis).Error; err != nil {
			level.Error(logger).Log("msg", "Failed to delete password that is too old", "err", err)
		}
	}
	return tx.Commit().Error
}

// UpdateLoginTime
//
//	@Description[en-US]: Update the user's last login time.
//	@Description[zh-CN]: 更新用户最后一次登陆时间。
//	@param ctx 	context.Context
//	@param id 	string
//	@return error
func (c *CommonService) UpdateLoginTime(ctx context.Context, id string) error {
	return c.Session(ctx).Model(&models.UserExt{UserId: id}).UpdateColumn("login_time", time.Now().UTC()).Error
}

type UserRole struct {
	UserId string `gorm:"primaryKey;column:user_id" json:"user_id"`
	RoleId string `gorm:"column:role_id" json:"role_id"`
	Role   string `gorm:"role" json:"role"`
}

func (c *CommonService) GetUsersRole(ctx context.Context, ids []string) (userRoles []*models.UserRole, err error) {
	var userRole []*UserRole
	if err = c.Session(ctx).Model(&models.UserRole{}).
		Select("`t_user_role`.`role_id`", "`t_user_role`.`user_id`", "`t_role`.`name` as `role`").
		Joins("JOIN `t_role` ON `t_role`.`id` = `t_user_role`.`role_id`").Find(&userRole, "user_id in ?", ids).Error; err != nil && err != gogorm.ErrRecordNotFound {
		return nil, err
	}
	for _, role := range userRole {
		userRoles = append(userRoles, (*models.UserRole)(role))
	}
	return userRoles, nil
}

func (c *CommonService) GetUserRole(ctx context.Context, id string) (role *models.Role, err error) {
	var userRole models.Role
	if err = c.Session(ctx).Model(&models.UserRole{}).Select("`t_role`.*").Joins("JOIN `t_role` ON `t_role`.`id` = `t_user_role`.`role_id`").First(&userRole, "user_id = ?", id).Error; err != nil && err != gogorm.ErrRecordNotFound {
		return nil, err
	}
	return &userRole, nil
}

func (c *CommonService) PatchUserRole(ctx context.Context, id string, roleName, roleId string) error {
	role := models.UserRole{UserId: id, RoleId: roleId}
	conn := c.Session(ctx)
	if len(roleId) == 0 {
		if len(roleName) == 0 {
			return conn.Delete(&role).Error
		}
		if err := c.Session(ctx).Select("id").Model(&models.Role{}).Where("name = ?", roleName).Scan(&roleId).Error; err != nil {
			return err
		}
	}

	return conn.Assign(models.UserRole{RoleId: roleId}).FirstOrCreate(&role).Error
}

func (c CommonService) InsertWeakPassword(ctx context.Context, passwords ...string) error {
	wps := make([]models.WeakPassword, len(passwords))
	for i, password := range passwords {
		wps[i] = models.WeakPassword{Hash: sign.SumSha512Hmac(password)}
	}
	conn := c.Session(ctx)
	return conn.Create(&wps).Error
}

func (c CommonService) VerifyWeakPassword(ctx context.Context, password string) error {
	logger := logs.GetContextLogger(ctx)
	hash := sign.SumSha512Hmac(password)

	conn := c.Session(ctx)
	var count int64
	if err := conn.Model(&models.WeakPassword{}).Where("hash = ?", hash).Count(&count).Error; err != nil {
		level.Error(logger).Log("msg", "Failed to check weak password", "err", err)
	}
	if count >= 1 {
		return errors.NewServerError(400, "Password too simple", errors.CodePasswordTooSimple)
	}
	return nil
}
