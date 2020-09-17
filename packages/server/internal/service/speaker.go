package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)



type Speaker struct {
	ID         string    `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name       string    `json:"name,omitempty"`
	Bio        string    `json:"bio,omitempty"`
	Headline   string    `json:"headline,omitempty"`
	Photo      string    `json:"photo,omitempty"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;"`
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

type CreateSpeakerInput struct {
	Name       string    `json:"name,omitempty"`
	Bio        string    `json:"bio,omitempty"`
	Headline   string    `json:"headline,omitempty"`
	Photo      string    `json:"photo,omitempty"`
}

func (api *api) CreateSpeaker(i CreateSpeakerInput) (*Speaker, error) {
	// The ID will be added on insert.
	speaker := Speaker{
		Name:     i.Name,
		Bio:      i.Bio,
		Headline: i.Headline,
		Photo:    i.Photo,
	}
	res := api.db.Create(&speaker)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &speaker, nil
}