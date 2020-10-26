package service

import (
	"campsite/packages/server/internal/service/role"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func (api *API) GetUserByEmail(email string) (*User, error)  {
	u := &User{Email: email}
	if err := api.db.First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (api *API) CreateUser(i SignUpInput) (*User, error)  {
	hash, err := api.generatePasswordHash(i.Password)
	if err != nil {
		return nil, err
	}
	u := User{
		ID:           uuid.New().String(),
		Email:        i.Email,
		PasswordHash: hash,
		Role:         role.USER,
	}
	if err := api.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (api *API) GetUserById(id string) (*User, error)  {
	u := User{ID: id}
	if err := api.db.First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (api *API) generatePasswordHash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), SALT_ROUND_COUNT)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

