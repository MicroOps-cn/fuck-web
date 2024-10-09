package gormservice

import (
	"context"
	"fmt"
	"strconv"

	gogorm "gorm.io/gorm"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck/clients/gorm"
)

func NewCommonService(name string, client *gorm.Client) *CommonService {
	return &CommonService{name: name, Client: client}
}

type CommonService struct {
	*gorm.Client
	name string
}

func (c CommonService) Name() string {
	return c.name
}

func (c CommonService) AutoMigrate(ctx context.Context) error {
	err := c.Session(ctx).AutoMigrate(
		&models.File{},
		&models.Permission{},
		&models.Role{},
		&models.AppKey{},
		&models.UserExt{},
		&models.UserRole{},
		&models.SystemConfig{},
		&models.UserPasswordHistory{},
		&models.WeakPassword{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c CommonService) RecordUploadFile(ctx context.Context, name string, path string, contentType string, size int64) (id string, err error) {
	file := &models.File{MimiType: contentType, Name: name, Path: path, Size: size}
	if err = c.Session(ctx).Create(file).Error; err != nil {
		return
	}
	return file.Id, err
}

func (c CommonService) GetFileInfoFromId(ctx context.Context, id string) (fileName, mimiType, filePath string, err error) {
	file := &models.File{Model: models.Model{Id: id}}
	if err = c.Session(ctx).First(file).Error; err != nil {
		return "", "", "", err
	}
	return file.Name, file.MimiType, file.Path, nil
}

func (c *CommonService) CreateTOTP(ctx context.Context, id string, secret string) error {
	tx := c.Session(ctx).Begin()
	defer tx.Rollback()
	ext := new(models.UserExt)
	if err := tx.Where("user_id = ?", id).First(&ext).Error; err == gogorm.ErrRecordNotFound {

		totp := models.UserExt{UserId: id, TOTPAsMFA: true}
		err = totp.SetSecret(secret)
		if err != nil {
			return err
		}
		if err = tx.Create(&totp).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	sec, err := ext.GetSecret()
	if err != nil || sec != secret {
		if err = ext.SetSecret(secret); err != nil {
			return err
		}
	}
	ext.TOTPAsMFA = true
	if err = tx.Where("user_id = ?", ext.UserId).
		Select("totp_salt", "totp_secret", "totp_as_mfa").Updates(ext).Error; err != nil {
		return errors.NewServerError(500, "failed to update totp setting: "+err.Error())
	}
	return tx.Commit().Error
}

func (c *CommonService) GetTOTPSecrets(ctx context.Context, ids []string) (secrets []string, err error) {
	conn := c.Session(ctx)
	var totps []models.UserExt
	err = conn.Where("user_id in ?", ids).Find(&totps).Error
	if err != nil {
		return nil, err
	}
	for _, totp := range totps {
		secret, err := totp.GetSecret()
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, secret)
	}
	return secrets, nil
}

func (c CommonService) PatchSystemConfig(ctx context.Context, prefix string, patch map[string]interface{}) error {
	tx := c.Session(ctx).Begin()
	defer tx.Rollback()
	for name, value := range patch {
		fullName := name
		if len(prefix) != 0 {
			fullName = fmt.Sprintf("%s.%s", prefix, name)
		}
		switch value.(type) {
		case string, uint, uint64, uint32, uint16, uint8, int, int64, int32, int16, int8, bool, float64, float32:
		default:
			return fmt.Errorf("unknown value type: %T", value)
		}
		valType := fmt.Sprintf("%T", value)
		var option models.SystemConfig
		if err := tx.Where("name = ?", fullName).First(&option).Error; err != nil {
			if err != gogorm.ErrRecordNotFound {
				return err
			} else if err = tx.Create(&models.SystemConfig{Name: fullName, Type: valType, Value: fmt.Sprintf("%v", value)}).Error; err != nil {
				return err
			}
			continue
		}
		if err := tx.Model(&models.SystemConfig{}).Where("name = ?", fullName).Updates(map[string]interface{}{
			"value": fmt.Sprintf("%v", value),
			"type":  valType,
		}).Error; err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

func (c CommonService) GetSystemConfig(ctx context.Context, prefix string) (map[string]interface{}, error) {
	conn := c.Session(ctx)
	var options []models.SystemConfig
	var count int64
	query := conn.Model(&models.SystemConfig{})
	if len(prefix) != 0 {
		query = conn.Where("name like ?", prefix+".%")
	}
	if err := query.Model(&models.SystemConfig{}).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 2000 {
		return nil, fmt.Errorf("There are too many configurations, please check. ")
	}
	if err := query.Limit(2000).Find(&options).Error; err != nil {
		return nil, err
	}

	cfgMap := map[string]interface{}{}
	for _, option := range options {
		name := option.Name[len(prefix)+1:]
		switch option.Type {
		case "string":
			cfgMap[name] = option.Value
		case "float64":
			if val, err := strconv.ParseFloat(option.Value, 64); err == nil {
				cfgMap[name] = val
			}
		case "float32":
			if val, err := strconv.ParseFloat(option.Value, 32); err == nil {
				cfgMap[name] = val
			}
		case "uint":
			if val, err := strconv.ParseUint(option.Value, 10, 32); err == nil {
				cfgMap[name] = uint(val)
			}
		case "uint64":
			if val, err := strconv.ParseUint(option.Value, 10, 64); err == nil {
				cfgMap[name] = val
			}
		case "uint32":
			if val, err := strconv.ParseUint(option.Value, 10, 32); err == nil {
				cfgMap[name] = uint32(val)
			}
		case "uint16":
			if val, err := strconv.ParseUint(option.Value, 10, 16); err == nil {
				cfgMap[name] = uint16(val)
			}
		case "uint8":
			if val, err := strconv.ParseUint(option.Value, 10, 8); err == nil {
				cfgMap[name] = uint8(val)
			}
		case "int":
			if val, err := strconv.ParseInt(option.Value, 10, 32); err == nil {
				cfgMap[name] = int(val)
			}
		case "int64":
			if val, err := strconv.ParseInt(option.Value, 10, 64); err == nil {
				cfgMap[name] = int64(val)
			}
		case "int32":
			if val, err := strconv.ParseInt(option.Value, 10, 32); err == nil {
				cfgMap[name] = int32(val)
			}
		case "int16":
			if val, err := strconv.ParseInt(option.Value, 10, 16); err == nil {
				cfgMap[name] = int16(val)
			}
		case "int8":
			if val, err := strconv.ParseInt(option.Value, 10, 8); err == nil {
				cfgMap[name] = int8(val)
			}
		case "bool":
			if val, err := strconv.ParseBool(option.Value); err == nil {
				cfgMap[name] = val
			}
		}
	}
	return cfgMap, nil
}
