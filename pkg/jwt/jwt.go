package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"net/http"
)

func VerifyToken(req *http.Request, jwtConfig *config.JwtConfig) (*model.Claims, error) {
	cookie, err := req.Cookie(jwtConfig.CookieName)
	if err != nil {
		return nil, err
	}

	tknString := cookie.Value

	claims := model.Claims{}
	tkn, err := jwt.ParseWithClaims(tknString, &claims, func(token *jwt.Token) (interface{}, error) {
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
