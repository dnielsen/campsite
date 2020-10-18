package service

import (
	"campsite/pkg/model"
	"github.com/google/uuid"
)

func (api *API) GetAllSpeakers() (*[]model.Speaker, error) {
	var speakers []model.Speaker
	if err := api.db.Find(&speakers).Error; err != nil {
		return nil, err
	}
	return &speakers, nil
}

func (api *API) GetSpeakerById(id string) (*model.Speaker, error) {
	s := model.Speaker{ID: id}
	if err := api.db.Preload("Sessions").First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (api *API) DeleteSpeakerById(id string) error {
	if err := api.db.Where("id = ?", id).Delete(&model.Speaker{}).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) CreateSpeaker(i model.SpeakerInput) (*model.Speaker, error) {
	s := model.Speaker{
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

func (api *API) EditSpeakerById(id string, i model.SpeakerInput) error {
	s := model.Speaker{
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
