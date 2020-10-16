package service

import "github.com/dgrijalva/jwt-go"


type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type SignUpInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

// Token will expire in 7 days from now.
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
