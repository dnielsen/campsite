package service

import "gorm.io/gorm"

type api struct {
	db *gorm.DB
}

func NewAPI(db *gorm.DB) *api {
	return &api{db}
}


type Datastore interface {
	GetSessionsByEventId(id string) (*[]Session, error)
	GetSessionById(id string) (*Session, error)
	GetAllSessions() (*[]Session, error)
	GetSessionsByIds(ids []string) (*[]Session, error)
}