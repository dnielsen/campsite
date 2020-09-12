package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const BASE_SESSION_API_URL = "http://localhost:5555"

type Session struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Title      string    `json:"title"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Description     string    `json:"description"`
}

func (api *api) GetSessionsByIds(ids []string) (*[]Session, error) {
	var body struct{
		SessionIds []string `json:"sessionIds"`
	}
	body.SessionIds = ids

	b, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal body: %v", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, BASE_SESSION_API_URL, bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.c.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
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
