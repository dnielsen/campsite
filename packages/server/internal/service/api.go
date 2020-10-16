package service

import (
	"campsite/packages/server/internal/config"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
	"os"
)

type API struct {
	db     *gorm.DB
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	return &API{db, c}
}

type EventAPI interface {
	AuthAPI
	GetAllEvents() (*[]Event, error)
	CreateEvent(userId string, i EventInput) (*Event, error)
	GetEventById(id string) (*Event, error)
	EditEventById(id string, i EventInput) error
	DeleteEventById(id string) error
}

type SessionAPI interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	EditSessionById(id string, i SessionInput) error
	DeleteSessionById(id string) error
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	EditSpeakerById(id string, i SpeakerInput) error
	DeleteSpeakerById(id string) error
}

type ImageAPI interface {
	GetImage(filename string) (*os.File, error)
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*Upload, error)
}

type AuthAPI interface {
	ValidateUser(i SignInInput) (*User, error)
	CreateUser(i SignUpInput) (*User, error)
	VerifyToken(r *http.Request) (*Claims, error)
	GenerateToken(email string) (string, error)
	GetUserByEmail(email string) (*User, error)
}