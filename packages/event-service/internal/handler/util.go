package handler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// We'll design it better later (move it to cookie/auth package).

const (
	ID                = "id"
	FILENAME          = "filename"
	TOKEN_HEADER_NAME = "Authorization"
	TOKEN_DURATION    = time.Hour * 24 * 7
)

// We'll later move it to an environment variable.
var JWT_SECRET_KEY = []byte("V3RY_S3CR3T_K3Y")

// Token will expire in 7 days from now.
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func verifyToken(r *http.Request) (*Claims, error) {
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

func generateToken(email string) (string, error) {
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
