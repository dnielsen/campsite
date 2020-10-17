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

func (api *API) CreateUser(i model.SignUpInput) (*model.User, error) {
	// Marshal the create user input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal create user input: %v", err)
		return nil, err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v", api.Config.Service.User.Host, api.Config.Service.User.Port), bytes.NewBuffer(b))
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
	var u model.User
	if err = json.Unmarshal(readBytes, &u); err != nil {
		log.Printf("Failed to unmarshal create user body: %v", err)
		return nil, err
	}
	return &u, nil
}

func (api *API) GetUserByEmail(email string) (*model.User, error) {
	// Create the request.
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", api.Config.Service.User.Host, api.Config.Service.User.Port, email), nil)
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
	var u model.User
	if err = json.Unmarshal(readBytes, &u); err != nil {
		log.Printf("Failed to unmarshal get user by id body: %v", err)
		return nil, err
	}
	return &u, nil
}