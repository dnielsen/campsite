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

func (api *API) SignIn(i model.SignInInput) (string, error) {
	// Marshal the sign in input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal sign in input: %v", err)
		return "", err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/sign-in", api.Config.Service.Auth.Host, api.Config.Service.Auth.Port), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return "", err
	}
	// Make the request.
	res, err := api.Client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return "", err
	}
	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return "", err
	}
	token := string(readBytes)
	return token, nil
}

func (api *API) SignUp(i model.SignUpInput) (string, error) {
	// Marshal the sign up input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal sign up input: %v", err)
		return "", err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/sign-up", api.Config.Service.Auth.Host, api.Config.Service.Auth.Port), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return "", err
	}
	// Make the request.
	res, err := api.Client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return "", err
	}
	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return "", err
	}
	token := string(readBytes)
	return token, nil
}


