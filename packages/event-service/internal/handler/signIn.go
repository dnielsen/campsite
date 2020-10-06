package handler

import (
	"campsite/packages/event-service/internal/service"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)



type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// We'll later move it to an environment variable.
var JWT_SECRET_KEY = []byte("V3RY_S3CR3T_K3Y")

func SignIn(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.SignInInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal sign in input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validate that the credentials match some user.
		u, err := api.ValidateUser(i)
		if err != nil {
			log.Printf("Failed to sign in: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Token will expire in 7 days from now.
		expiryTime := time.Now().Add(time.Hour * 24 * 7)

		claims := &Claims{
			Email:          u.Email,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed in Unix milliseconds.
				ExpiresAt: expiryTime.Unix(),
			},
		}

		// Sign the token.
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string.
		tokenString, err := token.SignedString(JWT_SECRET_KEY)
		if err != nil {
			log.Printf("Failed to create jwt token string: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the cookie header with the token for the client.
		http.SetCookie(w, &http.Cookie{
			Name:       TOKEN_COOKIE_NAME,
			Value:      tokenString,
			Expires:    expiryTime,
			// Currently we don't support `https` so we set
			// the `secure` property to `false`.
			Secure:     false,
			HttpOnly:   true,
			SameSite:   1,
		})

		// Respond that the sign in has been successful.
		w.WriteHeader(http.StatusOK)
	}
}