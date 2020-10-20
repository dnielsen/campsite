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
)

func (api *API) SignIn(i model.SignInInput) (string, error)  {
	// Validate the credentials match a user in the database.
	u, err := api.validateUser(i)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	// Get the JWT token.
	t, err := api.GenerateToken(u)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (api *API) validateUser(i model.SignInInput) (*model.User, error) {
	// Grab the user from the database.
	u, err := api.getUserByEmail(i.Email)
	if err != nil {
		return nil, err
	}
	// Verify the password is correct.
	if err := api.checkPasswordHash(u.PasswordHash, i.Password); err != nil {
		return nil, err
	}
	return u, nil
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
			ExpiresAt: time.Now().Add(TOKEN_DURATION).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tknStr, err := tkn.SignedString([]byte(api.c.Jwt.SecretKey))
	if err != nil {
		return "", err
	}
	return tknStr, nil
}

func (api *API) generatePasswordHash(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), SALT_ROUND_COUNT)
	if err != nil {
		return "", err
	}
	p := string(passwordBytes)
	return p, nil
}