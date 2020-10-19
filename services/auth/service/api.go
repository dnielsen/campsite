package service

import (
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"gorm.io/gorm"
)

type API struct {
	db     *gorm.DB
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	return &API{db, c}
}

type AuthAPI interface {
	ValidateUser(i model.SignInInput) (*model.User, error)
}
