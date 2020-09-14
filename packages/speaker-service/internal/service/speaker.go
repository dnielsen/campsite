package service

import (
	"github.com/lib/pq"
	"log"
)

type Speaker struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Name string    `json:"name"`
	Bio string `json:"bio"`
	Headline string `json:"headline"`
	Photo string `json:"photo"`
	SessionIds pq.StringArray `gorm:"type:uuid[]" json:"sessionIds"`
}

func (api *api) GetSpeakersByIds(ids []string) (*[]Speaker, error) {
	var speakers []Speaker
	_ = api.db.Where("id IN ?", ids).Find(&speakers)
	return &speakers, nil
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	_ = api.db.Find(&speakers)
	return &speakers, nil
}

func (api *api) GetSpeakerById(id string) (*Speaker, error) {
	var speaker Speaker
	_ = api.db.Where("id = ?", id).First(&speaker)
	log.Println(speaker.SessionIds)
	return &speaker, nil
}