package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dnielsen/campsite/pkg/model"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *API) CreateEvent(i model.EventInput) (*model.Event, error) {
	// Marshal the event input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal event input: %v", err)
		return nil, err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v", api.c.Service.Event.Host, api.c.Service.Event.Port), bytes.NewBuffer(b))
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
	var event model.Event
	if err = json.Unmarshal(readBytes, &event); err != nil {
		log.Printf("Failed to unmarshal event body: %v", err)
		return nil, err
	}
	return &event, nil
}

func (api *API) GetAllEvents() (*[]model.Event, error) {
	// Create the request.
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v", api.c.Service.Event.Host, api.c.Service.Event.Port), nil)
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
	var events []model.Event
	if err = json.Unmarshal(readBytes, &events); err != nil {
		log.Printf("Failed to unmarshal events body: %v", err)
		return nil, err
	}
	return &events, nil
}

func (api *API) GetEventById(id string) (*model.Event, error) {
	// Create the request.
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Event.Host, api.c.Service.Event.Port, id), nil)
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
	var event model.Event
	if err = json.Unmarshal(readBytes, &event); err != nil {
		log.Printf("Failed to unmarshal event body: %v", err)
		return nil, err
	}
	return &event, nil
}


func (api *API) EditEventById(id string, i model.EventInput) (*model.Event, error) {
	// Marshal the event input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal event input: %v", err)
		return nil, err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Event.Host, api.c.Service.Event.Port, id), bytes.NewBuffer(b))
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
	var event model.Event
	if err = json.Unmarshal(readBytes, &event); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}
	return &event, nil
}

func (api *API) DeleteEventById(id string) error {
	// Create the request.
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Session.Host, api.c.Service.Session.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return err
	}
	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return err
	}
	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return err
	}
	// Unmarshal the received body bytes.
	var event model.Event
	if err = json.Unmarshal(readBytes, &event); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return err
	}
	return nil
}
