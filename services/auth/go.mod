module github.com/dnielsen/campsite/services/auth

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dnielsen/campsite/pkg v0.0.0-20201030070358-3208b408ddc7
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx/v4 v4.9.1 // indirect
	github.com/rs/cors v1.7.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	golang.org/x/text v0.3.4 // indirect
	gorm.io/gorm v1.20.5
)
