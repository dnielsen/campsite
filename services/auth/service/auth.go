package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dnielsen/campsite/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	TOKEN_DURATION = time.Hour * 24 * 7
	SALT_ROUND_COUNT = 12
	TOKEN_HEADER_NAME = "Authorization"
	JWT_SECRET_KEY = "V3RY_S3CR3T_K3Y"
)


var USER_VALIDATION_ERR = errors.New("invalid credentials")

func (api *API) ValidateUser(i model.SignInInput) (*model.User, error) {
	u, err := api.getUserByEmail(i.Email)
	if err != nil {
		return nil, USER_VALIDATION_ERR
	}
	// Verify the password is correct.
	if err := api.checkPasswordHash
}

func (api *API) getUserByEmail(email string) (*model.User, error) {
	u := model.User{Email: email}
	if err := api.db.First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (api *API) checkPasswordHash(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

func (api *API) GenerateToken(u *model.User) (string, error) {
	claims := model.Claims{
		ID: u.ID,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TOKEN_DURATION),
			IssuedAt: time.Now(),
		},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tknStr, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tknStr, nil
}