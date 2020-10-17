package service

import (
	"campsite/pkg/config"
	"campsite/pkg/model"
	"gorm.io/gorm"
	"net/http"
)

type API model.API

func NewAPI(db *gorm.DB, c *config.Config) *API {
	// We define our own HttpClient to enable mocking (for easier testing).
	var client model.HttpClient = http.DefaultClient
	return &API{Db: db, Client: client, Config: c}
}
