package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Speaker struct {
	ID       string    `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name     string    `json:"name,omitempty"`
	Bio      string    `json:"bio,omitempty"`
	Headline string    `json:"headline,omitempty"`
	Photo    string    `json:"photo,omitempty"`
	Sessions []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;"`
}

func (api *api) GetSpeakerById(id string) (*Speaker, error) {
	var speaker Speaker
	res := api.db.Preload("Sessions").Where("id = ?", id).First(&speaker)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &speaker, nil
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	var speakers []Speaker
	res := api.db.Find(&speakers)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &speakers, nil
}

// Add UUID automatically on creation so that we can skip it in our methods
func (s *Speaker) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}

type SpeakerInput struct {
	// Name is a required field with a minimum and maximum length of 2 and 50 respectively.
	Name     string `json:"name,omitempty" validate:"required,min=2,max=50"`
	Bio      string `json:"bio,omitempty" validate:"required,min=20,max=2000"`
	Headline string `json:"headline,omitempty" validate:"required,min=2,max=30"`
	Photo    string `json:"photo,omitempty" validate:"required,min=10,max=150"`
}

func (api *api) CreateSpeaker(i SpeakerInput) (*Speaker, error) {
	// The ID will be added on insert.
	speaker := Speaker{
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

func (api *api) EditSpeaker(id string, i SpeakerInput) (*Speaker, error) {
	speakerUpdates := &Speaker{
		Name:     i.Name,
		Bio:      i.Bio,
		Headline: i.Headline,
		Photo:    i.Photo,
	}
	// Update the speaker in the database.
	if err := api.db.Model(&Speaker{}).Where("id = ?", id).Updates(&speakerUpdates).Error; err != nil {
		return nil, err
	}

	// Grab the updated speaker from the database.
	var s Speaker
	if err := api.db.Where("id = ?", id).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
