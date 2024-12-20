package gormservice

import (
	"context"
	"time"

	gogorm "gorm.io/gorm"

	"github.com/MicroOps-cn/fuck/clients/gorm"
	"github.com/MicroOps-cn/fuck/sets"
	w "github.com/MicroOps-cn/fuck/wrapper"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
)

func (c *CommonService) CreateRole(ctx context.Context, role *models.Role) (err error) {
	conn := c.Session(ctx)
	tx := conn.Begin()
	defer tx.Rollback()
	if err = tx.Create(role).Error; err != nil {
		return err
	}
	for _, permission := range role.Permission {
		if err = tx.Model(role).Association("Permission").Append(permission); err != nil {
			return errors.WithServerError(500, err, "failed to insert permission relationship")
		}
	}
	return tx.Commit().Error
}

func (c *CommonService) UpdateRole(ctx context.Context, role *models.Role) (err error) {
	conn := c.Session(ctx)
	return conn.Model(role).Association("Permission").Replace(role.Permission)
}

const sqlGetRolesPermission = `
SELECT 
    T0.id, 
    T0.name,  
    T1.role_id
FROM
    t_permission T0
        INNER JOIN
    t_rel_role_permission T1 ON T1.permission_id = T0.id
WHERE
    T1.role_id IN (?)
`

func (c *CommonService) GetRoles(ctx context.Context, keywords string, current, pageSize int64) (count int64, roles []*models.Role, err error) {
	conn := c.Session(ctx)
	tb := conn.Model(&models.Role{}).Where("delete_time is NULL")
	if keywords != "" {
		tb = tb.Where("Name LIKE ?", "%"+keywords+"%")
	}
	if err = tb.Count(&count).Error; err != nil {
		return 0, nil, err
	} else if count == 0 {
		return 0, nil, nil
	}
	if err = tb.Limit(int(pageSize)).Offset(int((current - 1) * pageSize)).Find(&roles).Error; err != nil {
		return 0, nil, err
	}
	var roleIds []interface{}
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}
	var permissions []models.RolePermission
	if err = conn.Raw(sqlGetRolesPermission, roleIds).Find(&permissions).Error; err == nil {
		for _, permission := range permissions {
			for idx, role := range roles {
				if role.Id == permission.RoleId {
					pm := permission.Permission
					roles[idx].Permission = append(roles[idx].Permission, &pm)
				}
			}
		}
	}
	return count, roles, nil
}

func (c CommonService) GetPermissions(ctx context.Context, keywords string, current int64, pageSize int64) (count int64, permissions []*models.Permission, err error) {
	conn := c.Session(ctx)
	tb := conn.Model(&models.Permission{})
	if keywords != "" {
		tb = tb.Where("Name LIKE ?", "%"+keywords+"%")
	}
	tb = tb.Where("enable_auth = 1")
	if err = tb.Count(&count).Error; err != nil {
		return 0, nil, err
	} else if count == 0 {
		return 0, nil, nil
	}
	if pageSize == 0 {
		pageSize = 1000
	}
	return count, permissions, tb.Limit(int(pageSize)).Offset(int((current - 1) * pageSize)).Find(&permissions).Error
}

func (c CommonService) DeleteRoles(ctx context.Context, ids []string) error {
	conn := c.Session(ctx)
	tx := conn.Begin()
	defer tx.Rollback()
	if err := tx.Model(models.Role{}).Where("id in (?)", ids).Association("Permission").Delete(); err != nil {
		return err
	} else if err = tx.Model(models.Role{}).Where("id in (?)", ids).Update("delete_time", time.Now().UTC()).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (c CommonService) CreateOrUpdateRoleByName(ctx context.Context, role *models.Role) error {
	newPermissions := role.Permission
	newPermissionNames := sets.New[string](w.Map(newPermissions, func(item *models.Permission) string {
		return item.Name
	})...)
	role.Permission = make([]*models.Permission, 0)
	conn := c.Session(ctx)
	if err := conn.Omit("Permission").FirstOrCreate(&role, "name = ?", role.Name).Error; err != nil {
		return err
	}
	if err := conn.Model(role).Association("Permission").Find(&role.Permission); err != nil {
		return err
	}
	oldPermissions := role.Permission
	oldPermissionNames := sets.New[string](w.Map(oldPermissions, func(item *models.Permission) string {
		return item.Name
	})...)
	needRemove := w.Map(oldPermissionNames.Difference(newPermissionNames).List(), func(name string) *models.Permission {
		return w.Find(oldPermissions, func(item *models.Permission) bool { return item.Name == name })
	})
	needInsert := w.Map(newPermissionNames.Difference(oldPermissionNames).List(), func(name string) *models.Permission {
		return w.Find(newPermissions, func(item *models.Permission) bool { return item.Name == name })
	})
	if err := conn.Model(role).Association("Permission").Append(&needInsert); err != nil {
		return err
	}
	if err := conn.Model(role).Association("Permission").Delete(&needRemove); err != nil {
		return err
	}
	role.Permission = append(role.Permission, needInsert...)
	return nil
}

func (c CommonService) RegisterPermission(ctx context.Context, permissions models.Permissions) error {
	conn := c.Session(ctx)
	var existsPermissions models.Permissions
	for i := 0; i < len(permissions); i += 10 {
		var names []string
		if len(permissions)-i < 50 {
			names = w.Map(permissions[i:], func(item *models.Permission) string { return item.Name })
		} else {
			names = w.Map(permissions[i:i+50], func(item *models.Permission) string { return item.Name })
		}
		if err := conn.Where("name in (?)", names).FindInBatches(&existsPermissions, 100, func(tx *gogorm.DB, batch int) error {
			var needUpdates models.Permissions
			for _, name := range names {
				p := w.Find(permissions, func(item *models.Permission) bool { return item.Name == name })
				op := w.Find(existsPermissions, func(item *models.Permission) bool { return item.Name == name })
				if op == nil {
					if err := conn.Create(p).Error; err != nil {
						return err
					}
				} else {
					p.Id = op.Id
					if op.EnableAuth != p.EnableAuth || op.Description != p.Description || op.ParentId != p.ParentId {
						op.EnableAuth = p.EnableAuth
						op.Description = p.Description
						op.ParentId = p.ParentId
						needUpdates = append(needUpdates, op)
					}
				}
				for _, child := range p.Children {
					child.ParentId = p.Id
				}

				if err := c.RegisterPermission(gorm.WithConnContext(ctx, conn), p.Children); err != nil {
					return err
				}
			}
			if len(needUpdates) > 0 {
				return tx.Save(needUpdates).Error
			}
			return nil
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

const sqlAuthorizationByRoleAndMethod = `
SELECT 
    COUNT(1) as count
FROM
    t_role T1
        JOIN
    t_rel_role_permission T2 ON T1.id = T2.role_id
        JOIN
    t_permission T3 ON T3.id = T2.permission_id
WHERE
    T1.name IN ? AND T3.name = ?
`

func (c CommonService) Authorization(ctx context.Context, roles []string, method string) bool {
	var count int64
	if err := c.Session(ctx).Raw(sqlAuthorizationByRoleAndMethod, roles, method).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}
