package service

import "time"

type Speaker struct {
	ID         string    `json:"id" gorm:"type:uuid"`
	Name       string    `json:"name" gorm:"not null"`
	Bio        string    `json:"bio" gorm:"not null"`
	Headline   string    `json:"headline" gorm:"not null"`
	Photo      string    `json:"photo" gorm:"not null"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SpeakerInput struct {
	Name string `json:"name"`
	Bio string `json:"bio"`
	Headline string `json:"headline"`
	Photo string `json:"photo"`
}


type Session struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `json:"name" gorm:"not null"`
	StartDate   *time.Time `json:"startDate" gorm:"not null"`
	EndDate     *time.Time `json:"endDate" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Event 		Event `json:"event,omitempty"`
	EventID 	string `json:"eventId,omitempty" gorm:"type:uuid;not null"`
	Speakers    []Speaker `json:"speakers,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
	Comments []Comment `json:"comments,omitempty"`
}

type Event struct {
	ID            string         `json:"id" gorm:"type:uuid"`
	Name          string         `json:"name" gorm:"not null"`
	RegistrationUrl string `json:"registrationUrl" gorm:"not null"`
	Description   string         `json:"description" gorm:"not null"`
	StartDate     *time.Time      `json:"startDate" gorm:"not null"`
	EndDate       *time.Time      `json:"endDate" gorm:"not null"`
	Photo         string         `json:"photo" gorm:"not null"`
	OrganizerName string         `json:"organizerName" gorm:"not null"`
	Address       *string  `json:"address"`
	Sessions      []Session      `json:"sessions,omitempty"`
}

type Comment struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Session 		Session `json:"session,omitempty"`
	SessionID 	string `json:"sessionId,omitempty" gorm:"type:uuid;not null"`
}