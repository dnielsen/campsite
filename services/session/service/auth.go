package service

import (
	"github.com/dnielsen/campsite/pkg/jwt"
	"github.com/dnielsen/campsite/pkg/model"
	"net/http"
)

func (api *API) VerifyToken(r *http.Request) (*model.Claims, error) {
	claims, err := jwt.VerifyToken(r, &api.c.Jwt)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
