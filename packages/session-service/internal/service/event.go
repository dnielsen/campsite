package service

import "time"

//type Event struct {
//	ID            string         `json:"id"`
//	Name          string         `json:"name"`
//	Description   string         `json:"description"`
//	StartDate     *time.Time      `json:"startDate"`
//	EndDate       *time.Time      `json:"endDate"`
//	Photo         string         `json:"photo"`
//	OrganizerName string         `json:"organizerName"`
//	Address       *string  `json:"address"`
//	Sessions      []Session      `json:"sessions,omitempty"`
//	Speakers      []Speaker      `json:"speakers,omitempty"`
//}

type Event struct {
	ID            string         `json:"id" gorm:"type:uuid"`
	Name          string         `json:"name" gorm:"not null"`
	Description   string         `json:"description" gorm:"not null"`
	StartDate     *time.Time      `json:"startDate" gorm:"not null"`
	EndDate       *time.Time      `json:"endDate" gorm:"not null"`
	Photo         string         `json:"photo" gorm:"not null"`
	OrganizerName string         `json:"organizerName" gorm:"not null"`
	Address       *string  `json:"address"`
	Sessions      []Session      `json:"sessions,omitempty"`
}


