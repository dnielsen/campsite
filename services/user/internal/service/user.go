package service

import (
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

func (api *API) CreateUser(i CreateUserInput) (*User, error)  {
	// Hash the password so that we store it in our database encrypted.
	hash, err := generatePasswordHash(i.Password)
	if err != nil {
		return nil, err
	}
	// Create the user in the database.
	u := User{
		ID:           uuid.New().String(),
		Email:        i.Email,
		PasswordHash: hash,
	}
	if err := api.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (api *API) GetUserById(id string) (*User, error) {
	u := User{ID: id}
	if err := api.db.First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}