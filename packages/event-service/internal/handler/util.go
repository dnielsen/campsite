package handler

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// We'll design it better later (move it to cookie/auth package).

const (
	ID = "id"
	FILENAME = "filename"
	TOKEN_COOKIE_NAME = "Authorization"
)

// Token will expire in 7 days from now.
var EXPIRY_TIME = time.Now().Add(time.Hour * 24 * 7)


type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func verifyToken(w http.ResponseWriter, r *http.Request) (*Claims, error) {
	tokenString := r.Header.Get(TOKEN_COOKIE_NAME)
	claims := Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	// Clear the cookie when token is invalid.
	if !tkn.Valid {
		http.SetCookie(w, nil)
	}

	return &claims, nil
}

func generateToken(email string) (string, error) {
	claims := Claims{
		Email:          email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed in Unix milliseconds.
			ExpiresAt: EXPIRY_TIME.Unix(),
		},
	}

	// Sign the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	// Create the JWT string.
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}


	return tokenString, nil
}

func setTokenCookie(w http.ResponseWriter, token string) {
	// Set the cookie header with the token for the client.
	c := http.Cookie{
		Name:    TOKEN_COOKIE_NAME,
		Value:   token,
		Expires: EXPIRY_TIME,
		// Currently we don't support `https` so we set
		// the `secure` property to `false`.
		Secure:     false,
		HttpOnly:   true,
		// No same site since our frontend url differs from the API url right now.
		SameSite:   1,
	}

	//w.Header().Set(CO, c.String())
	http.SetCookie(w, &c)

}