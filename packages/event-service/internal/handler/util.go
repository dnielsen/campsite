package handler

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

const (
	ID = "id"
	FILENAME = "filename"
	// We'll design it better later (move it).
	TOKEN_COOKIE_NAME = "token"
)


func verifyToken(w http.ResponseWriter, r *http.Request) (*Claims, error) {
	c, err := r.Cookie(TOKEN_COOKIE_NAME)
	if err != nil {
		return nil, err
	}
	tokenString := c.Value
	claims := Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	// Clear the cookie when token is invalid.
	if !tkn.Valid {
		http.SetCookie(w, nil)
	}

	return &claims, nil
}
