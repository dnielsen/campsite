package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

type SessionInput struct {
	Name        string     `json:"name,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
	Url         string     `json:"url,omitempty"`
	SpeakerIds []string `json:"speakerIds,omitempty"`
	EventId string `json:"eventId,omitempty"`
}

func (api *API) GetAllSessions() (*[]Session, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v", api.c.Service.Session.Host, api.c.Service.Session.Port), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return nil, err
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes.
	var sessions []Session
	if err = json.Unmarshal(readBytes, &sessions); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}

	return &sessions, nil
}

func (api *API) GetSessionById(id string) (*Session, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Session.Host, api.c.Service.Session.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return nil, err
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes.
	var session Session
	if err = json.Unmarshal(readBytes, &session); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}

	return &session, nil
}

func (api *API) DeleteSessionById(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Session.Host, api.c.Service.Session.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return err
	}

	// Make the request.
	if _, err := api.client.Do(req); err != nil {
		log.Printf("Failed to do request: %v", err)
		return err
	}

	return nil
}

func (api *API) EditSessionById(id string, i SessionInput) error {
	// Marshal the session input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal session input: %v", err)
		return err
	}

	// Create a request.
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Session.Host, api.c.Service.Session.Port, id), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return err
	}

	// Make the request.
	if _, err := api.client.Do(req); err != nil {
		log.Printf("Failed to do request: %v", err)
		return err
	}

	return nil
}

func (api *API) CreateSession(i SessionInput) (*Session, error) {
	// Marshal the session input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal session input: %v", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v", api.c.Service.Session.Host, api.c.Service.Session.Port), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return nil, err
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes.
	var session Session
	if err = json.Unmarshal(readBytes, &session); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}

	return &session, nil
}