package service

type CreateUserInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Role string `json:"role"`
}