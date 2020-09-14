package service

import (
	"github.com/lib/pq"
	"time"
)

type Session struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Description string    `json:"description"`
	SpeakerIds pq.StringArray `json:"speakerIds" gorm:"type:uuid[]"`
	EventId string 		   `gorm:"type:uuid"`
}

func (api *api) GetSessionsByIds(ids []string) (*[]Session, error) {
	var sessions []Session
	_ = api.db.Where("id IN ?", ids).Find(&sessions)
	return &sessions, nil
}

func (api *api) GetSessionsByEventId(id string) (*[]Session, error) {
	var sessions []Session
	_ = api.db.Where("event_id = ?", id).Find(&sessions)
	return &sessions, nil
}

func (api *api) GetSessionById(id string) (*Session, error) {
	var session Session
	_ = api.db.Where("id = ?", id).First(&session)
	return &session, nil
}

func (api *api) GetAllSessions() (*[]Session, error) {
	var sessions []Session
	_ = api.db.Find(&sessions)
	return &sessions, nil
}