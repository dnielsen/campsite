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

type CommentAPI interface {
	CreateComment(sessionId string, i model.CommentInput) (*model.Comment, error)
	GetCommentsBySessionId(sessionId string, limit string, cursor string) (*model.CommentResponse, error)
}

type AuthAPI interface {
	SignIn(i model.SignInInput) (string, error)
	SignUp(i model.SignUpInput) (string, error)
}