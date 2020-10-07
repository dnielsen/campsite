package service

import "time"

type Event struct {
	ID              string     `json:"id" gorm:"type:uuid"`
	Name            string     `json:"name" gorm:"not null"`
	RegistrationUrl string     `json:"registrationUrl" gorm:"not null"`
	Description     string     `json:"description" gorm:"not null"`
	StartDate       *time.Time `json:"startDate" gorm:"not null"`
	EndDate         *time.Time `json:"endDate" gorm:"not null"`
	Photo           string     `json:"photo" gorm:"not null"`
	OrganizerName   string     `json:"organizerName" gorm:"not null"`
	Address         *string    `json:"address"`
	Sessions        []Session  `json:"sessions,omitempty"`
}
