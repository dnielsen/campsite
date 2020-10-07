package service

import (
	"github.com/google/uuid"
)

type Speaker struct {
	ID       string    `json:"id" gorm:"type:uuid"`
	Name     string    `json:"name" gorm:"not null"`
	Bio      string    `json:"bio" gorm:"not null"`
	Headline string    `json:"headline" gorm:"not null"`
	Photo    string    `json:"photo" gorm:"not null"`
	Sessions []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SpeakerInput struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Headline string `json:"headline"`
	Photo    string `json:"photo"`
}

func (api *API) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	if err := api.db.Find(&speakers).Error; err != nil {
		return nil, err
	}
	return &speakers, nil
}

func (api *API) GetSpeakerById(id string) (*Speaker, error) {
	s := Speaker{ID: id}
	if err := api.db.Preload("Sessions").First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (api *API) DeleteSpeakerById(id string) error {
	if err := api.db.Where("id = ?", id).Delete(&Speaker{}).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) CreateSpeaker(i SpeakerInput) (*Speaker, error) {
	s := Speaker{
		ID:       uuid.New().String(),
		Name:     i.Name,
		Bio:      i.Bio,
		Headline: i.Headline,
		Photo:    i.Photo,
	}
	if err := api.db.Create(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (api *API) EditSpeakerById(id string, i SpeakerInput) error {
	s := Speaker{
		ID:       id,
		Name:     i.Name,
		Bio:      i.Bio,
		Headline: i.Headline,
		Photo:    i.Photo,
	}
	if err := api.db.Updates(&s).Error; err != nil {
		return err
	}
	return nil
}
