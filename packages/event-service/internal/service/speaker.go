package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BASE_SPEAKER_API_URL = "http://localhost:3333"

type Speaker struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Name string    `json:"name"`
	Bio string `json:"bio"`
	Headline string `json:"headline"`
	Photo string `json:"photo"`
}

func (api *api) GetSpeakersByIds(ids []string) (*[]Speaker, error) {
	var body struct{
		SpeakerIds []string `json:"speakerIds"`
	}
	body.SpeakerIds = ids

	b, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal body: %v", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/byIds",BASE_SPEAKER_API_URL), bytes.NewBuffer(b))
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

	// Unmarshal the received body bytes
	var speakers []Speaker
	if err = json.Unmarshal(readBytes, &speakers); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speakers, nil
}


func (api *api) GetSpeakerById(id string) (*Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v",BASE_SPEAKER_API_URL, id), nil)
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

	// Unmarshal the received body bytes
	var speaker Speaker
	if err = json.Unmarshal(readBytes, &speaker); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speaker, nil
}


func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_SPEAKER_API_URL, nil)
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

	// Unmarshal the received body bytes
	var speakers []Speaker
	if err = json.Unmarshal(readBytes, &speakers); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speakers, nil
}
