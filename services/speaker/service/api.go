package service

import (
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"gorm.io/gorm"
)

type API struct {
	db *gorm.DB
	c *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	return &API{db, c}
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]model.Speaker, error)
	CreateSpeaker(i model.SpeakerInput) (*model.Speaker, error)
	GetSpeakerById(id string) (*model.Speaker, error)
	EditSpeakerById(id string, i model.SpeakerInput) (*model.Speaker, error)
	DeleteSpeakerById(id string) error
}

