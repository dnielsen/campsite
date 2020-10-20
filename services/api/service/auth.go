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

func (api *API) SignIn(i model.SignInInput) (string, error) {
	// Marshal the sign in input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal sign in input: %v", err)
		return "", err
	}
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/sign-in", api.c.Service.Auth.Host, api.c.Service.Auth.Port), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return "", err
	}
	// Make the request.
	res, err := api.client.Do(req)
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
