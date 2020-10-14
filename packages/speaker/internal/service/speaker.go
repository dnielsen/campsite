package service

import (
	"github.com/google/uuid"
)

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
