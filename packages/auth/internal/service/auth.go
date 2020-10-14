package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GITHUB_USER_API_URL = "https://api.github.com/user"
	// This should be an environment valuable
	CLIENT_ID = "df486625c9c9cce6ccd7"
	// This should be an environment valuable
	CLIENT_SECRET = "4984affda80501a590dee1d7cbb84a56aa08ca82"
	GITHUB_OAUTH_BASE_URL = "https://github.com/login/oauth"
)

func (api *API) GetGitHubAccessToken(code string) (string, error) {
	body := OAuthTokenInput{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		Code:         code,
	}

	// Marshal the body.
	b, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal body: %v", err)
		return "", err
	}

	// Construct the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/access_token", GITHUB_OAUTH_BASE_URL), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Make the request
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return "", err
	}

	// Read the received body stream.
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read body: %v", err)
		return "", nil
	}

	log.Printf("%v", string(resBody))
	// Unmarshal the received body bytes.
	var token OAuthTokenResponse
	if err := json.Unmarshal(resBody, &token); err != nil {
		log.Printf("Failed to unmarshal OAuth token response body: %v", err)
		return "", nil
	}

	return token.AccessToken, nil
}

func (api *API) GetGitHubUserData(token string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, GITHUB_USER_API_URL, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return "", err
	}

	return string(bodyBytes), nil
}

// TODO: implement me
func (api *API) ValidateUser(i SignInInput) (*User, error) {
	return &User{}, nil
}

func (api *API) VerifyGitHubToken(r *http.Request) (*User, error)  {
	//req, err := http.NewRequest(http.MethodGet, GITHUB_USER_API_URL, nil)
	//if err != nil {
	//	return nil, err
	//}
	//// Set access token query parameter
	//q := url.Values{}
	//q.Add("Authorization", fmt.Sprintf("token %s"))
	panic("das")
}