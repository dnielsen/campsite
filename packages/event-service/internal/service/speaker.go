package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Speaker struct {
	ID         string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name       string    `json:"name"`
	Bio        string    `json:"bio"`
	Headline   string    `json:"headline"`
	Photo      string    `json:"photo"`
	SessionIds []string  `json:"sessionIds"`
	Sessions   []Session `json:"sessions,omitempty"`
}

func (api *api) GetSpeakersByIds(ids []string) (*[]Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v?%v", api.c.Service.Speaker.Address, strings.Join(ids, ",")), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes
	var speakers []Speaker
	if err = json.Unmarshal(readBytes, &speakers); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speakers, nil
}

func (api *api) GetSpeakerById(id string) (*Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v", api.c.Service.Speaker.Address, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes
	var speaker Speaker
	if err = json.Unmarshal(readBytes, &speaker); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speaker, nil
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, api.c.Service.Speaker.Address, nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes
	var speakers []Speaker
	if err = json.Unmarshal(readBytes, &speakers); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speakers, nil
}
