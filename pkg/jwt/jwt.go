package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"net/http"
)

const TOKEN_HEADER_NAME = "Authorization"

func VerifyToken(req *http.Request, jwtConfig *config.JwtConfig) (*model.Claims, error) {
	tokenString := req.Header.Get(TOKEN_HEADER_NAME)
	claims := model.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtConfig.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("token invalid")
	}

	return &claims, nil
}
