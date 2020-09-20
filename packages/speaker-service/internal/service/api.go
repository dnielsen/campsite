package service

import "gorm.io/gorm"

type api struct {
	db *gorm.DB
}



func NewAPI(db *gorm.DB) *api {
	return &api{db}
}

type Datastore interface {
	GetSpeakersByIds(ids []string) (*[]Speaker, error)
	GetAllSpeakers() (*[]Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	DeleteSpeakerById(id string) error
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
}