package service

import (
	"github.com/google/uuid"
)

type Speaker struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	Name string `json:"name"`
}

type SpeakerDatastore interface {
	GetAllSpeakers() (*[]Speaker, error)
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	_ = api.db.Find(&speakers)
	return &speakers, nil
}