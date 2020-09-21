package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//type Speaker struct {
//	ID        string `json:"id"`
//	Name string    `json:"name"`
//	Bio string `json:"bio"`
//	Headline string `json:"headline"`
//	Photo string `json:"photo"`
//	Sessions []Session `json:"sessions,omitempty"`
//}

type Speaker struct {
	ID         string    `json:"id" gorm:"type:uuid"`
	Name       string    `json:"name" gorm:"not null"`
	Bio        string    `json:"bio" gorm:"not null"`
	Headline   string    `json:"headline" gorm:"not null"`
	Photo      string    `json:"photo" gorm:"not null"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;"`
}


func (api *api) GetSpeakersByIds(ids []string) (*[]Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v?id=%v", api.c.Service.Speaker.Address, strings.Join(ids, ",")), nil)
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
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received bytes.
	var speakers []Speaker
	if err = json.Unmarshal(bytes, &speakers); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speakers, nil
}
