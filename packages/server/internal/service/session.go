package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Session struct {
	ID          string     `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name        string     `json:"name,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
	Url         string     `json:"url,omitempty"`
	Event       *Event     `json:"event,omitempty"`
	EventID     *string     `json:"-"`
	Speakers    []Speaker  `json:"speakers,omitempty" gorm:"many2many:session_speakers;"`
}

func (api *api) GetSessionById(id string) (*Session, error) {
	var session Session
	res := api.db.Preload("Speakers").Where("id = ?", id).First(&session)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *api) GetAllSessions() (*[]Session, error) {
	var sessions []Session
	res := api.db.Find(&sessions)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &sessions, nil
}

// Add UUID automatically on creation so that we can skip it in our methods
func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}

type SessionInput struct {
	Name        string     `json:"name,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
	Url         string     `json:"url,omitempty"`
	SpeakerIds []string `json:"speakerIds,omitempty"`
}

func (api *api) CreateSession(i SessionInput) (*Session, error) {
	// Get the speakers from the database to attach them to the session.
	var speakers []Speaker
	if err := api.db.Where("id IN ?", i.SpeakerIds).Find(&speakers).Error; err != nil {
		return nil, err
	}
	// Create the session with the speakers attached.
	session := Session{
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		Url:         i.Url,
		Speakers: speakers,
	}
	if err := api.db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *api) EditSession(id string, i SessionInput) (*Session, error) {
	sessionUpdates := &Session{
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		Url:         i.Url,
	}
	// Update the session in the database.
	if err := api.db.Model(&Session{}).Where("id = ?", id).Updates(&sessionUpdates).Error; err != nil {
		return nil, err
	}

	// Grab the updated session from the database.
	var s Session
	if err := api.db.Where("id = ?", id).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
