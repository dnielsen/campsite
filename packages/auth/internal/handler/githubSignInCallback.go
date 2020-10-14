package handler

import (
	"campsite/packages/auth/internal/service"
	"log"
	"net/http"
)


type OAuthTokenInput struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code string `json:"code"`
}

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
}

func GitHubSignInCallback(api service.AuthAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get(CODE)
		t, err := api.GetGitHubAccessToken(code)
		if err != nil {
			log.Printf("Failed to get GitHub access token: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := api.GetGitHubUserData(t)
		if err != nil {
			log.Printf("Failed to get GitHub user data: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}


		log.Printf("github user data: %v", data)
	}
}
