package gormservice

import (
	"context"

	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
	"github.com/MicroOps-cn/fuck/clients/gorm"
)

func NewUserAndAppService(_ context.Context, name string, client *gorm.Client) *UserService {
	set := &UserService{name: name, Client: client}
	return set
}

type UserService struct {
	*gorm.Client
	name string
}

func (s UserService) AutoMigrate(ctx context.Context) error {
	err := s.Session(ctx).AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}
