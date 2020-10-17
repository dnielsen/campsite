package service

import (
	"campsite/pkg/model"
	"github.com/google/uuid"
)

func (api *API) GetSessionById(id string) (*model.Session, error) {
	s := model.Session{ID: id}
	if err := api.Db.Preload("Speakers").Preload("Event.Sessions").First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (api *API) GetAllSessions() (*[]model.Session, error) {
	var ss []model.Session
	if err := api.Db.Find(&ss).Error; err != nil {
		return nil, err
	}
	return &ss, nil
}

func (api *API) DeleteSessionById(id string) error {
	if err := api.Db.Where("id = ?", id).Delete(&model.Session{}).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) CreateSession(i model.SessionInput) (*model.Session, error) {
	var spks []model.Speaker
	if err := api.Db.Where("id IN ?", i.SpeakerIds).Find(&spks).Error; err != nil {
		return nil, err
	}
	session := model.Session{
		ID:          uuid.New().String(),
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		EventID:     i.EventId,
		Url:         i.Url,
		Speakers:    spks,
	}
	if err := api.Db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *API) EditSessionById(id string, i model.SessionInput) (*model.Session, error) {
	// Update the session (besides speakers).
	sess := model.Session{
		ID:          id,
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		Url:         i.Url,
		EventID:     i.EventId,
	}
	if err := api.Db.Updates(&sess).Error; err != nil {
		return nil, err
	}
	// Update the session speakers. We're doing it this way instead of adding it to the
	// session struct because otherwise we would just add new speaker ids to the session_speaker table
	// instead of replacing them.

	// Get the speakers with just their ids.
	var spks []model.Speaker
	if err := api.Db.Where("id IN ?", i.SpeakerIds).Select("id").Find(&spks).Error; err != nil {
		return &sess, err
	}
	// Update the speakers
	if err := api.Db.Model(&sess).Association("Speakers").Replace(spks); err != nil {
		return &sess, err
	}
	// Return the updated session.
	return &sess, nil
}
