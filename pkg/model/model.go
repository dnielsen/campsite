package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Speaker struct {
	ID       string    `json:"id,omitempty" gorm:"type:uuid"`
	Name     string    `json:"name,omitempty" gorm:"not null"`
	Bio      string    `json:"bio,omitempty" gorm:"not null"`
	Headline string    `json:"headline,omitempty" gorm:"not null"`
	Photo    string    `json:"photo,omitempty" gorm:"not null"`
	Sessions []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SpeakerInput struct {
	Name     string `json:"name,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Headline string `json:"headline,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

type Event struct {
	ID              string     `json:"id,omitempty" gorm:"type:uuid"`
	Name            string     `json:"name,omitempty" gorm:"not null"`
	Description     string     `json:"description,omitempty" gorm:"not null"`
	RegistrationUrl string     `json:"registrationUrl,omitempty" gorm:"not null"`
	StartDate       *time.Time `json:"startDate,omitempty" gorm:"not null"`
	EndDate         *time.Time `json:"endDate,omitempty" gorm:"not null"`
	Photo           string     `json:"photo,omitempty" gorm:"not null"`
	OrganizerName   string     `json:"organizerName,omitempty" gorm:"not null"`
	Address         *string    `json:"address,omitempty"`
	Sessions        []Session  `json:"sessions,omitempty"`
	Speakers        []Speaker  `json:"speakers,omitempty" gorm:"-"`
}

type EventInput struct {
	Name            string     `json:"name,omitempty"`
	StartDate       *time.Time `json:"startDate,omitempty"`
	EndDate         *time.Time `json:"endDate,omitempty"`
	RegistrationUrl string     `json:"registrationUrl,omitempty"`
	Description     string     `json:"description,omitempty"`
	Photo           string     `json:"photo,omitempty"`
	OrganizerName   string     `json:"organizerName,omitempty"`
	Address         *string    `json:"address,omitempty"`
}

type Session struct {
	ID          string     `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name        string     `json:"name,omitempty" gorm:"not null"`
	StartDate   *time.Time `json:"startDate,omitempty" gorm:"not null"`
	EndDate     *time.Time `json:"endDate,omitempty" gorm:"not null"`
	Description string     `json:"description,omitempty" gorm:"not null"`
	VideoUrl    string     `json:"url,omitempty" gorm:"not null"`
	Event       Event      `json:"event,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	EventID     string     `json:"eventId,omitempty" gorm:"type:uuid;not null"`
	Speakers    []Speaker  `json:"speakers,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SessionInput struct {
	Name        string     `json:"name,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
	VideoUrl    string     `json:"videoUrl,omitempty"`
	SpeakerIds  []string   `json:"speakerIds,omitempty"`
	EventId     string     `json:"eventId,omitempty"`
}

type User struct {
	ID string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	// We likely don't ever wanna expose the password hash.
	PasswordHash string `json:"-"`
}

type SignInInput struct {
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Claims struct {
	Me
	jwt.StandardClaims
}

type Upload struct {
	Url string `json:"url,omitempty"`
}

type Me struct {
	ID string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}