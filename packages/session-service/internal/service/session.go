package service

import (
	"time"
)

type Session struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Title      string    `json:"title"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Description     string    `json:"description"`
}

type SessionDatastore interface {
	GetSessionsByIds(ids []string) (*[]Session, error)
}

func (api *api) GetSessionsByIds(ids []string) (*[]Session, error) {
	var sessions []Session
	_ = api.db.Find(&sessions, ids)
	return &sessions, nil
}
