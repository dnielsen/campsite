package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type CommentResponse struct {
	Comments *[]Comment `json:"comments"`
	EndCursor *string `json:"endCursor"`
}

func (api *API) GetCommentsBySessionId(sessionId string, limit string, cursor string) (*CommentResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v/comments", api.c.Service.Session.Host, api.c.Service.Session.Port, sessionId), nil)
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
	var commentResponse CommentResponse
	if err = json.Unmarshal(readBytes, &commentResponse); err != nil {
		log.Printf("Failed to unmarshal comments body: %v", err)
		return nil, err
	}

	return &commentResponse, nil
}