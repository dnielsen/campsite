package service

import (
	"gorm.io/gorm"
)

type api struct {
	db     *gorm.DB
}

func NewAPI(db *gorm.DB) *api {
	return &api{db}
}

type EventDatastore interface {
	GetEventById(id string) (*Event, error)
}