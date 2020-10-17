package service

import (
	"campsite/packages/server/internal/service/role"
	"github.com/dgrijalva/jwt-go"
	"time"
)



type Session struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `json:"name" gorm:"not null"`
	StartDate   *time.Time `json:"startDate" gorm:"not null"`
	EndDate     *time.Time `json:"endDate" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Event 		Event `json:"event,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	EventID 	string `json:"eventId,omitempty" gorm:"type:uuid;not null"`
	Speakers    []Speaker `json:"speakers,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SessionInput struct {
	Name        string     `json:"name,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
	Url         string     `json:"url,omitempty"`
	SpeakerIds []string `json:"speakerIds,omitempty"`
	EventId string `json:"eventId,omitempty"`
}


type Upload struct {
	Url string `json:"url"`
}

type Event struct {
	ID            string     `json:"id" gorm:"type:uuid"`
	Name          string     `json:"name" gorm:"not null"`
	Description   string     `json:"description" gorm:"not null"`
	RegistrationUrl string `json:"registrationUrl" gorm:"not null"`
	StartDate     *time.Time `json:"startDate" gorm:"not null"`
	EndDate       *time.Time `json:"endDate" gorm:"not null"`
	Photo         string     `json:"photo" gorm:"not null"`
	OrganizerName string     `json:"organizerName" gorm:"not null"`
	Address       *string    `json:"address"`
	Sessions      []Session  `json:"sessions"`
	Speakers []Speaker `json:"speakers,omitempty" gorm:"-"`
	User 		User `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	UserID 	string `json:"userId,omitempty" gorm:"type:uuid;not null"`

}

type EventInput struct {
	Name string `json:"name,omitempty"`
	StartDate     *time.Time `json:"startDate,omitempty"`
	EndDate       *time.Time `json:"endDate,omitempty"`
	RegistrationUrl string `json:"registrationUrl,omitempty"`
	Description   string     `json:"description,omitempty"`
	Photo         string     `json:"photo,omitempty"`
	OrganizerName string     `json:"organizerName,omitempty"`
	Address       *string    `json:"address,omitempty"`
}

type Speaker struct {
	ID         string    `json:"id" gorm:"type:uuid"`
	Name       string    `json:"name" gorm:"not null"`
	Bio        string    `json:"bio" gorm:"not null"`
	Headline   string    `json:"headline" gorm:"not null"`
	Photo      string    `json:"photo" gorm:"not null"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SpeakerInput struct {
	// Name is a required field with a minimum and maximum length of 2 and 50 respectively.
	Name     string `json:"name,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Headline string `json:"headline,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID string `json:"id" gorm:"type:uuid;not null"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Events []Event `json:"events"`
	Role role.Role `json:"role"`
}

// Token will expire in 7 days from now.
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
