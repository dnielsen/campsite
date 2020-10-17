package service

import (
	"campsite/packages/session/internal/config"
	"gorm.io/gorm"
	"net/http"
)

type API struct {
	db     *gorm.DB
	client HttpClient
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	// We're using our custom `HttpClient` to enable mocking.
	var client HttpClient = http.DefaultClient

	return &API{db, client, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SessionAPI interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	EditSessionById(id string, i SessionInput) error
	DeleteSessionById(id string) error
}

type CommentAPI interface {
	CreateComment(sessionId string, i CommentInput) (*Comment, error)
	GetCommentsBySessionId(sessionId string, limit int, cursor string) (*[]Comment, *string, error)
}
