package service

import "github.com/google/uuid"

type Speaker struct {
	ID         string    `json:"id" gorm:"type:uuid"`
	Name       string    `json:"name" gorm:"not null"`
	Bio        string    `json:"bio" gorm:"not null"`
	Headline   string    `json:"headline" gorm:"not null"`
	Photo      string    `json:"photo" gorm:"not null"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SpeakerInput struct {
	// Name is a required field with a minimum and maximum length of 2 and 50 respectively.
	Name     string `json:"name,omitempty" validate:"required,min=2,max=50"`
	Bio      string `json:"bio,omitempty" validate:"required,min=20,max=2000"`
	Headline string `json:"headline,omitempty" validate:"required,min=2,max=30"`
	Photo    string `json:"photo,omitempty" validate:"required,min=10,max=150"`
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	if err := api.db.Find(&speakers).Error; err != nil {
		return nil, err
	}
	return &speakers, nil
}

func (api *api) CreateSpeaker(i SpeakerInput) (*Speaker, error) {
	speaker := Speaker{
		ID:         uuid.New().String(),
		Name:       i.Name,
		Bio:        i.Bio,
		Headline:   i.Headline,
		Photo:      i.Photo,
	}
	if err := api.db.Create(&speaker).Error; err != nil {
		return nil, err
	}
	return &speaker, nil
}

func (api *api) GetSpeakerById(id string) (*Speaker, error) {
	speaker := Speaker{ID: id}
	if err := api.db.Preload("Sessions").First(&speaker).Error; err != nil {
		return nil, err
	}
	return &speaker, nil
}

func (api *api) DeleteSpeakerById(id string) error {
	speaker := Speaker{ID: id}
	if err := api.db.Delete(&speaker).Error; err != nil {
		return err
	}
	return nil
}