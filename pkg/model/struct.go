package model

import (
	"campsite/pkg/config"
	"campsite/pkg/role"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"
)

type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type SignUpInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Role role.Role `json:"role"`
}

// Token will expire in 7 days from now.
type Claims struct {
	ID string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type API struct {
	Db     *gorm.DB
	Client HttpClient
	Config *config.Config
}


type Speaker struct {
	ID       string    `json:"id" gorm:"type:uuid"`
	Name     string    `json:"name" gorm:"not null"`
	Bio      string    `json:"bio" gorm:"not null"`
	Headline string    `json:"headline" gorm:"not null"`
	Photo    string    `json:"photo" gorm:"not null"`
	Sessions []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}


type SpeakerInput struct {
	Name     string `json:"name,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Headline string `json:"headline,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

type Event struct {
	ID              string     `json:"id" gorm:"type:uuid"`
	Name            string     `json:"name" gorm:"not null"`
	Description     string     `json:"description" gorm:"not null"`
	RegistrationUrl string     `json:"registrationUrl" gorm:"not null"`
	StartDate       *time.Time `json:"startDate" gorm:"not null"`
	EndDate         *time.Time `json:"endDate" gorm:"not null"`
	Photo           string     `json:"photo" gorm:"not null"`
	OrganizerName   string     `json:"organizerName" gorm:"not null"`
	Address         *string    `json:"address"`
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
	ID          string     `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string     `json:"name" gorm:"not null"`
	StartDate   *time.Time `json:"startDate" gorm:"not null"`
	EndDate     *time.Time `json:"endDate" gorm:"not null"`
	Description string     `json:"description" gorm:"not null"`
	Url         string     `json:"url" gorm:"not null"`
	Event       Event      `json:"event,omitempty"`
	EventID     string     `json:"eventId,omitempty" gorm:"type:uuid;not null"`
	Speakers    []Speaker  `json:"speakers,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
	Comments    []Comment  `json:"comments,omitempty"`
}

type SessionInput struct {
	Name        string     `json:"name,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
	Url         string     `json:"url,omitempty"`
	SpeakerIds  []string   `json:"speakerIds,omitempty"`
	EventId     string     `json:"eventId,omitempty"`
}

type Comment struct {
	ID        string    `gorm:"primaryKey;type:uuid" json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	SessionID string    `json:"sessionId,omitempty" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE;"`
}

type CommentInput struct {
	// We don't need the session id since we get it from the url param.
	Content string `json:"content"`
}

type CommentResponse struct {
	Comments  *[]Comment `json:"comments"`
	EndCursor *string    `json:"endCursor"`
}

type Upload struct {
	Url string `json:"url"`
}
