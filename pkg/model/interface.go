package model

import (
	"mime/multipart"
	"net/http"
	"os"
)

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type UserAPI interface {
	AuthAPI
	CreateUser(i SignUpInput) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type AuthAPI interface {
	ValidateUser(i SignInInput) (*User, error)
	VerifyToken(r *http.Request) (*Claims, error)
	GenerateToken(user *User) (string, error)
	GetUserByEmail(email string) (*User, error)
}

type EventAPI interface {
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	GetEventById(id string) (*Event, error)
	EditEventById(id string, i EventInput) (*Event, error)
	DeleteEventById(id string) error
}

type SessionAPI interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	EditSessionById(id string, i SessionInput) (*Session, error)
	DeleteSessionById(id string) error
	CreateComment(sessionId string, i CommentInput) (*Comment, error)
	GetCommentsBySessionId(sessionId string, limit string, cursor string) (*CommentResponse, error)
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	EditSpeakerById(id string, i SpeakerInput) (*Speaker, error)
	DeleteSpeakerById(id string) error
}

type ImageAPI interface {
	GetImage(filename string) (*os.File, error)
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*Upload, error)
}
