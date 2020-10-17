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
	// Create the request.
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v", api.Config.Service.Session.Host, api.Config.Service.Session.Port), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}
	// Make the request.
	res, err := api.Client.Do(req)
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
	var ss []model.Session
	if err = json.Unmarshal(readBytes, &ss); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}
	return &ss, nil
}

func (api *API) GetSessionById(id string) (*model.Session, error) {
	// Create the request.
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", api.Config.Service.Session.Host, api.Config.Service.Session.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}
	// Make the request.
	res, err := api.Client.Do(req)
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
	var s model.Session
	if err = json.Unmarshal(readBytes, &s); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}
	return &s, nil
}

func (api *API) DeleteSessionById(id string) error {
	// Create the request.
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%v:%v/%v", api.Config.Service.Session.Host, api.Config.Service.Session.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return err
	}
	// Make the request.
	if _, err := api.Client.Do(req); err != nil {
		log.Printf("Failed to do request: %v", err)
		return err
	}
	return nil
}

func (api *API) EditSessionById(id string, i model.SessionInput) (*model.Session, error) {
	// Marshal the session input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal session input: %v", err)
		return nil, err
	}
	// Create a request.
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%v:%v/%v", api.Config.Service.Session.Host, api.Config.Service.Session.Port, id), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}
	// Make the request.
	res, err := api.Client.Do(req);
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
	var s model.Session
	if err = json.Unmarshal(readBytes, &s); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}
	return &s, nil
}

func (api *API) CreateSession(i model.SessionInput) (*model.Session, error) {
	// Marshal the session input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal session input: %v", err)
		return nil, err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v", api.Config.Service.Session.Host, api.Config.Service.Session.Port), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}
	// Make the request.
	res, err := api.Client.Do(req)
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
	var s model.Session
	if err = json.Unmarshal(readBytes, &s); err != nil {
		log.Printf("Failed to unmarshal session body: %v", err)
		return nil, err
	}
	return &s, nil
}

func (api *API) CreateComment(sessionId string, i model.CommentInput) (*model.Comment, error) {
	// Marshal the comment input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal comment input: %v", err)
		return nil, err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/%v/comments", api.Config.Service.Session.Host, api.Config.Service.Session.Port, sessionId), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}
	// Make the request.
	res, err := api.Client.Do(req)
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
	var c model.Comment
	if err = json.Unmarshal(readBytes, &c); err != nil {
		log.Printf("Failed to unmarshal comment body: %v", err)
		return nil, err
	}
	return &c, nil
}
