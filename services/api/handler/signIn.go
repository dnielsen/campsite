package handler

import (
	"github.com/dnielsen/campsite/services/api/service"
	"net/http"
)

// `/auth/sign-in` POST route. It communicates with the event service only.
func SignIn(api service.AuthAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api.SignIn(w, r)
	}
}