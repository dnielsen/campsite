package handler

import (
	"fmt"
	"net/http"
)

const (
	CLIENT_ID     = "CLIENT_ID"
	CLIENT_SECRET = "CLIENT_SECRET"
)

func githubSignIn(w http.ResponseWriter, r *http.Request) {
	callbackRedirectUri := "http://localhost:4444/sign-in/github/callback"
	redirectUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", CLIENT_ID, callbackRedirectUri)

	http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
}
