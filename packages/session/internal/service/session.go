package service

import (
	"github.com/google/uuid"
)

func (api *API) GetSessionById(id string) (*Session, error) {
	s := Session{ID: id}
	if err := api.db.Preload("Speakers").Preload("Event.Sessions").First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (api *API) GetAllSessions() (*[]Session, error) {
	var ss []Session
	if err := api.db.Find(&ss).Error; err != nil {
		return nil, err
	}
	return &ss, nil
}

func (api *API) DeleteSessionById(id string) error {
	if err := api.db.Where("id = ?", id).Delete(&Session{}).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) CreateSession(i SessionInput) (*Session, error) {
	var spks []Speaker
	if err := api.db.Where("id IN ?", i.SpeakerIds).Find(&spks).Error; err != nil {
		return nil, err
	}
	session := Session{
		ID:          uuid.New().String(),
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		EventID:     i.EventId,
		Url:         i.Url,
		Speakers: spks,
	}
	if err := api.db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *API) EditSessionById(id string, i SessionInput) error {
	// Update the session (besides speakers).
	sess := Session{
		ID:          id,
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		Url:         i.Url,
		EventID:     i.EventId,
	}
	if err := api.db.Updates(&sess).Error; err != nil {
		return err
	}
	// Update the session speakers. We're doing it this way instead of adding it to the
	// session struct because otherwise we would just add new speaker ids to the session_speaker table
	// instead of replacing them.

	// Get the speakers with just their ids.
	var spks []Speaker
	if err := api.db.Where("id IN ?", i.SpeakerIds).Select("id").Find(&spks).Error; err != nil {
		return err
	}
	// Update the speakers
	if err := api.db.Model(&sess).Association("Speakers").Replace(spks); err != nil {
		return err
	}
	return nil
}

