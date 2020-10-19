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
	SignIn(i model.SignInInput) (string, error)
	getUserByEmail(email string) (*model.User, error)
	validateUser(i model.SignInInput) (*model.User, error)
	checkPasswordHash(passwordHash string, password string) error
	generatePasswordHash(password string) (string, error)
}
