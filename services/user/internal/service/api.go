package service

import (
	"campsite/services/user/internal/config"
	"gorm.io/gorm"
	"net/http"
)

type API struct {
	db     *gorm.DB
	client HttpClient
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	// We define our own HttpClient to enable mocking (for easier testing).
	var client HttpClient = http.DefaultClient

	return &API{db, client, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type UserAPI interface {
	CreateUser(i CreateUserInput) (*User, error)
	GetUserById(id string) (*User, error)
}
