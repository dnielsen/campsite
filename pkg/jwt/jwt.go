package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
)

func VerifyToken(tokenString string, jwtConfig *config.JwtConfig) (*model.Claims, error) {
	claims := model.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("token invalid")
	}

	return &claims, nil
}
