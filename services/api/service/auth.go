package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *API) SignIn(w http.ResponseWriter, r *http.Request) {
	// Create the request.
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/sign-in", api.c.Service.Auth.Host, api.c.Service.Auth.Port), r.Body)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := string(readBytes)
	// Send the cookie with the token.
	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      token,
		Path:       "/",
		MaxAge:     60 * 60 * 24 * 7,
		Secure:     false,
		HttpOnly:   true,
		SameSite:   0,
	})
	// Respond with the received status code.
	w.WriteHeader(res.StatusCode)
}
