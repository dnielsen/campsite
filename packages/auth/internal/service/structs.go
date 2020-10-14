package service

type OAuthTokenInput struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code string `json:"code"`
}

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
}

type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Email string `json:"email"`
}