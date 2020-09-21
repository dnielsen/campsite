package service

import "time"


type Session struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `json:"name" gorm:"not null"`
	StartDate   *time.Time `json:"startDate" gorm:"not null"`
	EndDate     *time.Time `json:"endDate" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Event 		Event `json:"-"`
	EventID 	string `json:"-" gorm:"type:uuid;not null"`
	Speakers    []Speaker `json:"speakers,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

