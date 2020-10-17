package service

import "github.com/dgrijalva/jwt-go"


type CreateUserInput struct {
	SignInInput
}

type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type SignUpInput struct {
	CreateUserInput
}


type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

// Token will expire in 7 days from now.
type Claims struct {
	ID string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
