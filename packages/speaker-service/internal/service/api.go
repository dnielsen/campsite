package service

import "gorm.io/gorm"

type API struct {
	db *gorm.DB
}

func NewAPI(db *gorm.DB) *API {
	return &API{db}
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	EditSpeakerById(id string, i SpeakerInput) error
	DeleteSpeakerById(id string) error
}