package service

import (
	"bytes"
	"campsite/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *API) GetAllSessions() (*[]model.Session, error) {
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
	var sessions []model.Session
	if err = json.Unmarshal(readBytes, &sessions); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}

	return &sessions, nil
}

func (api *API) GetSessionById(id string) (*model.Session, error) {
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
	var session model.Session
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

func (api *API) EditSessionById(id string, i model.SessionInput) error {
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

func (api *API) CreateSession(i model.SessionInput) (*model.Session, error) {
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
	var session model.Session
	if err = json.Unmarshal(readBytes, &session); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}

	return &session, nil
}
