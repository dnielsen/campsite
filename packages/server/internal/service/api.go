package service

import (
	"campsite/packages/server/internal/config"
	"campsite/packages/server/internal/service/role"
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
	CreateEvent(i EventInput, userId string) (*Event, error)
	GetEventById(id string) (*Event, error)
	EditEventById(id string, i EventInput) error
	DeleteEventById(id string) error
}

type SessionAPI interface {
	AuthAPI
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput, userId string) (*Session, error)
	GetSessionById(id string) (*Session, error)
	EditSessionById(id string, i SessionInput) error
	DeleteSessionById(id string) error
}

type SpeakerAPI interface {
	AuthAPI
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput, userId string) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	EditSpeakerById(id string, i SpeakerInput) error
	DeleteSpeakerById(id string) error
}

type ImageAPI interface {
	AuthAPI
	GetImage(filename string) (*os.File, error)
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*Upload, error)
}

type AuthAPI interface {
	ValidateUser(i SignInInput) (*User, error)
	VerifyToken(r *http.Request) (*Claims, error)
	GenerateToken(user *User) (string, error)
	VerifyRole(userId string, roles []role.Role) (*User, error)
}

type UserAPI interface {
	AuthAPI
	CreateUser(i SignUpInput) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
}