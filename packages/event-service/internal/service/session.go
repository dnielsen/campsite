package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const BASE_SESSION_API_URL = "http://localhost:5555"

type Session struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Description string    `json:"description"`
	Speakers []Speaker `json:"speakers,omitempty"`
	SpeakerIds []string `json:"-"`
}

func (api *api) GetAllSessions() (*[]Session, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_SESSION_API_URL, nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.c.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}
	log.Println(res)

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

func (api *api) GetSessionsByEventId(id string) (*[]Session, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/event/%v",BASE_SESSION_API_URL, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.c.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}
	log.Println(res)

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

func (api *api) GetSessionById(id string) (*Session, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v", BASE_SESSION_API_URL, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.c.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}
	log.Println(res)

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

