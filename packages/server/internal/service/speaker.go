package service

import "github.com/google/uuid"

type Speaker struct {
	ID       string    `json:"id" gorm:"type:uuid"`
	Name     string    `json:"name" gorm:"not null"`
	Bio      string    `json:"bio" gorm:"not null"`
	Headline string    `json:"headline" gorm:"not null"`
	Photo    string    `json:"photo" gorm:"not null"`
	Sessions []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SpeakerInput struct {
	// Name is a required field with a minimum and maximum length of 2 and 50 respectively.
	Name     string `json:"name,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Headline string `json:"headline,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

func (api *API) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	if err := api.db.Find(&speakers).Error; err != nil {
		return nil, err
	}
	return &speakers, nil
}

func (api *API) CreateSpeaker(i SpeakerInput) (*Speaker, error) {
	speaker := Speaker{
		ID:       uuid.New().String(),
		Name:     i.Name,
		Bio:      i.Bio,
		Headline: i.Headline,
		Photo:    i.Photo,
	}
	if err := api.db.Create(&speaker).Error; err != nil {
		return nil, err
	}
	return &speaker, nil
}

func (api *API) GetSpeakerById(id string) (*Speaker, error) {
	speaker := Speaker{ID: id}
	// We're preloading sessions since we need them in the speaker by id page.
	// For now we're getting all of the properties, we'll optimize it later.
	if err := api.db.Preload("Sessions").First(&speaker).Error; err != nil {
		return nil, err
	}
	return &speaker, nil
}

func (api *API) DeleteSpeakerById(id string) error {
	speaker := Speaker{ID: id}
	if err := api.db.Delete(&speaker).Error; err != nil {
		return err
	}
	return nil
}
