package service

import (
	"campsite/pkg/model"
	"gorm.io/gorm"
)

type API struct {
	db *gorm.DB
}

func NewAPI(db *gorm.DB) *API {
	return &API{db}
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]model.Speaker, error)
	CreateSpeaker(i model.SpeakerInput) (*model.Speaker, error)
	GetSpeakerById(id string) (*model.Speaker, error)
	EditSpeakerById(id string, i model.SpeakerInput) error
	DeleteSpeakerById(id string) error
}
