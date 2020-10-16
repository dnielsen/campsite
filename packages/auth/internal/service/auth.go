package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	TOKEN_HEADER_NAME = "Authorization"
	TOKEN_DURATION    = time.Hour * 24 * 7
	SALT_ROUND_COUNT = 12
)

// We'll later move it to an environment variable.
var JWT_SECRET_KEY = []byte("V3RY_S3CR3T_K3Y")

// TODO: implement me
func (api *API) ValidateUser(i SignInInput) (*User, error) {
	u, err := api.GetUserByEmail(i.Email)
	if err != nil {
		return nil, err
	}
	if isPasswordValid := api.checkPasswordHash(u.PasswordHash, i.Password); !isPasswordValid{
		return nil, err
	}
	return u, nil
}

func (api *API) checkPasswordHash(passwordHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func (api *API) generatePasswordHash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), SALT_ROUND_COUNT)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

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
	}
	if err := api.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (api *API) VerifyToken(r *http.Request) (*Claims, error) {
	tokenString := r.Header.Get(TOKEN_HEADER_NAME)
	claims := Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}


func (api *API) GenerateToken(email string) (string, error) {
	claims := Claims{
		Email:          email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed in Unix milliseconds.
			ExpiresAt: time.Now().Add(TOKEN_DURATION).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}

	// Sign the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	// Create the JWT string.
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}