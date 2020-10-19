package service

import (
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"mime/multipart"
	"net/http"
	"os"
)

type API struct {
	client HttpClient
	c      *config.Config
}

func NewAPI(c *config.Config) *API {
	// We define our own HttpClient to enable mocking (for easier testing).
	// We don't have tests yet, however, it's a common practice to do that
	// for this reason.
	var client HttpClient = http.DefaultClient
	return &API{client, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type EventAPI interface {
	GetAllEvents() (*[]model.Event, error)
	CreateEvent(i model.EventInput) (*model.Event, error)
	GetEventById(id string) (*model.Event, error)
	EditEventById(id string, i model.EventInput) error
	DeleteEventById(id string) error
}

type SessionAPI interface {
	GetAllSessions() (*[]model.Session, error)
	CreateSession(i model.SessionInput) (*model.Session, error)
	GetSessionById(id string) (*model.Session, error)
	EditSessionById(id string, i model.SessionInput) error
	DeleteSessionById(id string) error
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]model.Speaker, error)
	CreateSpeaker(i model.SpeakerInput) (*model.Speaker, error)
	GetSpeakerById(id string) (*model.Speaker, error)
	EditSpeakerById(id string, i model.SpeakerInput) error
	DeleteSpeakerById(id string) error
}

type ImageAPI interface {
	GetImage(filename string) (*os.File, error)
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*model.Upload, error)
}
