package main

import (
	"campsite/pkg/config"
	"campsite/pkg/database"
	"campsite/pkg/middleware"
	"campsite/services/auth/handler"
	"campsite/services/auth/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

const (
	SERVICE_NAME = "auth"
	READ_TIMEOUT  = 15 * time.Second
	WRITE_TIMEOUT = 15 * time.Second
	IDLE_TIMEOUT  = 120 * time.Second
)

func main() {

	// Initialize the config which includes
	// database, server, and other services' configuration.
	c := config.NewConfig()

	// Create a new database connection. Also, since it's a dev db,
	// migrate it and create sample mock data there.
	db := database.NewDb(&c.Db)

	// Set up the API.
	api := service.NewAPI(db, c)

	// Set up the router.
	r := mux.NewRouter()

	// Logger middleware logs the incoming requests.
	// Example output: `status=200 method=GET path=/events duration=3.714408ms`
	r.Use(middleware.Logger)
	// Request limiter middleware ensures that a client
	// with a given IP can only make so many requests.
	// If he does too many within a certain amount of time,
	// the server responds with the status code 429 (Too Many Requests).
	// This middleware is only used in the Event Service,
	// since this service is the only service that can call
	// all the others.
	r.Use(middleware.RequestLimiter)

	// Enable tracing - forward our requests to the zipkin server.
	if c.Tracing.Enabled == true {
		r.Use(middleware.Tracing("auth", c.Service.Auth.Port, c))
		log.Println("Tracing middleware has been enabled")
	}

	// Register our handlers.
	// Plain JWT sign in
	r.HandleFunc("/sign-in", handler.SignIn(api)).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", handler.SignUp(api)).Methods(http.MethodPost)

	// For dev only - Set up CORS so our client (React app) can consume the api.
	corsWrapper := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "Authorization", "Cookie", "token", "*"},
	})

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Service.Auth.Port),
		Handler:      corsWrapper.Handler(r),
		ReadTimeout:  READ_TIMEOUT,
		WriteTimeout: WRITE_TIMEOUT,
		IdleTimeout:  IDLE_TIMEOUT,
	}

	// startServer the server.
	log.Printf("Listening at: %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
