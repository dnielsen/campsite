package service

import (
	"time"
)

type Session struct {
	ID          string    `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name        string    `json:"name,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	EndDate     time.Time `json:"endDate,omitempty"`
	Description string    `json:"description,omitempty"`
	Url         string    `json:"url,omitempty"`
	Event		Event 	   `json:"event,omitempty"`
	EventID     string     `json:"-"`
	Speakers    []Speaker `json:"speakers,omitempty" gorm:"many2many:session_speakers;"`
}
