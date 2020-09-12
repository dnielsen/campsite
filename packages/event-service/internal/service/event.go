package service

import (
	"github.com/lib/pq"
	"time"
)

type Event struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string    `json:"name"`
	Description string `json:"description"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Photo     string    `json:"photo"`
	// Temporarily there's no Organizer structure
	OrganizerName string   `json:"organizerName"`
	Address       string   `json:"address"`
	Sessions []Session `json:"sessions,omitempty" gorm:"-"`
	SpeakerIds    pq.StringArray `gorm:"type:uuid[]" json:"-"`
	Speakers []Speaker `json:"speakers,omitempty" gorm:"-"`
}

func (api *api) GetEventById(id string) (*Event, error) {
	var event Event
	// TODO
	// The line below doesn't work for some reason that is it results in:
	// SELECT * FROM "events" WHERE 82911d7b-6884-4f9f-bd32-e88ac6e0d952 ORDER BY "events"."id" LIMIT 1
	// `id = ` is missing for some reason. Docs say it should work.
	//
	// _ = api.db.First(&event, id)
	_ = api.db.Where("id = ?", id).First(&event)
	return &event, nil
}