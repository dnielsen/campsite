package service

import (
	"campsite/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)


func (api *API) GetCommentsBySessionId(sessionId string, limit string, cursor string) (*model.CommentResponse, error) {
	// Create the request.
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v/comments", api.Config.Service.Session.Host, api.Config.Service.Session.Port, sessionId), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}
	// Set the query parameters.
	q := url.Values{}
	q.Add("limit", limit)
	q.Add("cursor", cursor)
	// Encode the query parameters.
	req.URL.RawQuery = q.Encode()
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
	var commentResponse model.CommentResponse
	if err = json.Unmarshal(readBytes, &commentResponse); err != nil {
		log.Printf("Failed to unmarshal comments body: %v", err)
		return nil, err
	}
	return &commentResponse, nil
}
