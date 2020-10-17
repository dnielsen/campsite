package service

import (
	"campsite/pkg/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const SALT_ROUND_COUNT = 12

func generatePasswordHash(password string) (string, error)  {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), SALT_ROUND_COUNT)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (api *API) CreateUser(i model.SignUpInput) (*model.User, error)  {
	// Hash the password so that we store it in our database encrypted.
	hash, err := generatePasswordHash(i.Password)
	if err != nil {
		return nil, err
	}
	// Create the user in the database.
	u := model.User{
		ID:           uuid.New().String(),
		Email:        i.Email,
		PasswordHash: hash,
	}
	if err := api.Db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (api *API) GetUserById(id string) (*model.User, error) {
	u := model.User{ID: id}
	if err := api.Db.First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}