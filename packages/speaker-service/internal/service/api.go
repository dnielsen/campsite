package service

import "gorm.io/gorm"

type api struct {
	db *gorm.DB
}



func NewAPI(db *gorm.DB) *api {
	return &api{db}
}

type Datastore interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	DeleteSpeakerById(id string) error
	GetSpeakersByIds(ids []string) (*[]Speaker, error)
}