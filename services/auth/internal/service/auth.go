package service

import (
	"campsite/packages/auth/internal/service/role"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	TOKEN_HEADER_NAME = "Authorization"
	TOKEN_DURATION    = time.Hour * 24 * 7
	SALT_ROUND_COUNT = 12
)

// We'll later move it to an environment variable.
var JWT_SECRET_KEY = []byte("V3RY_S3CR3T_K3Y")

func (api *API) ValidateUser(i SignInInput) (*User, error) {
	// We need `validationErr` so that whenever email or password don't match,
	// it says the same message, that is we avoid giving out the information to the client
	// about registered emails
	validationErr := errors.New("invalid credentials")
	// Grab the user from the database.
	u, err := api.GetUserByEmail(i.Email)
	if err != nil {
		return nil, validationErr
	}
	// Verify the password is correct.
	if err := api.checkPasswordHash(u.PasswordHash, i.Password); err != nil {
		return nil, validationErr
	}
	return u, nil
}

func (api *API) checkPasswordHash(passwordHash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

func (api *API) VerifyToken(r *http.Request) (*Claims, error) {
	tokenString := r.Header.Get(TOKEN_HEADER_NAME)
	claims := Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}


// GenerateToken returns a jwt token string and an error if the token is somehow invalid
func (api *API) GenerateToken(u *User) (string, error) {
	claims := Claims{
		ID: u.ID,
		Email:          u.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed in Unix milliseconds.
			ExpiresAt: time.Now().Add(TOKEN_DURATION).Unix(),
			IssuedAt: time.Now().Unix(),
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

// Checks if the user role is equal to one of the specified roles.
func (api *API) VerifyRole(u *User, roleWhitelist []role.Role) (*User, error) {
	// If user role isn't contained in the roles array then the user doesn't have
	// the permissions needed.
	if isPermitted := role.Contains(roleWhitelist, u.Role); !isPermitted {
		// We're returning the user anyways, it might be useful in the future.
		return u, errors.New("no permissions")
	}
	return u, nil
}