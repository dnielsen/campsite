package main

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/database"
	"github.com/dnielsen/campsite/pkg/middleware"
	"github.com/dnielsen/campsite/services/auth/handler"
	"github.com/dnielsen/campsite/services/auth/service"
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

	// Create a new database connection.
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

	// Enable tracing middleware - forward our request data to the zipkin server
	// that is running with Hypertrace.
	if c.Tracing.Enabled == true {
		r.Use(middleware.Tracing(SERVICE_NAME, c.Service.Auth.Port, c))
		log.Println("Tracing middleware has been enabled")
	}

	// Register our handlers.
	r.HandleFunc("/auth/sign-in", handler.SignIn(api)).Methods(http.MethodPost)

	// Set up the server.
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Service.Auth.Port),
		Handler:      corsWrapper.Handler(r),
		ReadTimeout:  READ_TIMEOUT,
		WriteTimeout: WRITE_TIMEOUT,
		IdleTimeout:  IDLE_TIMEOUT,
	}

	// Start the server.
	log.Printf("Listening at: %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
