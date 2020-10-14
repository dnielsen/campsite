package handler

import (
	"campsite/packages/auth/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func GitHubSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//redirectUrl := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s", service.GITHUB_OAUTH_BASE_URL, service.CLIENT_ID, GITHUB_CALLBACK_REDIRECT_URI)
		//http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
		code := r.URL.Query().Get(CODE)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/access_token", service.GITHUB_OAUTH_BASE_URL), nil)
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set query parameters
		q := url.Values{}
		q.Add("client_id", service.CLIENT_ID)
		q.Add("client_secret", service.CLIENT_SECRET)
		q.Add("code", code)

		req.URL.RawQuery = q.Encode()
		req.Header.Set("accept", "application/json")

		//res, err := api.
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Failed to do request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var t OAuthTokenResponse
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			log.Printf("Failed to decode OAuth token response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", "http://localhost:3000")
		w.WriteHeader(http.StatusFound)
		http.SetCookie(w, &http.Cookie{
			Name:       "token",
			Value:      t.AccessToken,
			Expires:    time.Now().Add(time.Hour * 24 * 7),
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
		})
		w.Write([]byte(t.AccessToken))
	}
}