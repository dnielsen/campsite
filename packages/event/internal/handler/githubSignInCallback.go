package handler
//
//import (
//	"bytes"
//	"campsite/packages/event/internal/service"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//const (
//	CODE = "code"
//	GITHUB_OAUTH_BASE_URL = "https://github.com/login/oauth"
//)
//
//type OAuthTokenInput struct {
//	ClientID string `json:"client_id"`
//	ClientSecret string `json:"client_secret"`
//	Code string `json:"code"`
//}
//
//type OAuthTokenResponse struct {
//	AccessToken string `json:"access_token"`
//	TokenType string `json:"token_type"`
//	Scope string `json:"scope"`
//}
//
//
//func (api *service.API) getGithubAccessToken(code string) (string, error) {
//	body := OAuthTokenInput{
//		ClientID:     CLIENT_ID,
//		ClientSecret: CLIENT_SECRET,
//		Code:         CODE,
//	}
//
//	b, err := json.Marshal(body)
//	if err != nil {
//		log.Printf("Failed to marshal body: %v", err)
//		return "", err
//	}
//	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/access_token", GITHUB_OAUTH_BASE_URL), bytes.NewBuffer(b))
//	if err != nil {
//		log.Printf("Failed to create request: %v", err)
//		return "", err
//	}
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Accept", "application/json")
//
//	res, err := api.c.Do(req)
//	if err != nil {
//		log.Printf("Failed to do request: %v", err)
//		return "", err
//	}
//
//	resBody, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		log.Printf("Failed to read body: %v", err)
//		return "", nil
//	}
//	var token OAuthTokenResponse
//	if err := json.Unmarshal(resBody, &token); err != nil {
//		log.Printf("Failed to unmarshal OAuth token response body: %v", err)
//		return "", nil
//	}
//
//	return token.AccessToken, nil
//}
//
//
//func githubSignInCallback(w http.ResponseWriter, r *http.Request) {
//	code := r.URL.Query().Get(CODE)
//	githubAccessToken :=
//
//		gi
//}
