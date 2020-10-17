package service

import (
	"campsite/pkg/config"
	"campsite/pkg/model"
	"gorm.io/gorm"
	"net/http"
)

type API model.API

func NewAPI(db *gorm.DB, c *config.Config) *API {
	// We're using our custom `HttpClient` to enable mocking.
	var client model.HttpClient = http.DefaultClient
	return &API{Db: db, Client: client, Config: c}
}