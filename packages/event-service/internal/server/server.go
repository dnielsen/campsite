package server

import (
	"campsite/packages/event-service/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

const (
	READ_TIMEOUT  = 15 * time.Second
	WRITE_TIMEOUT = 15 * time.Second
	IDLE_TIMEOUT  = 120 * time.Second
)

func Start(r *mux.Router, c *config.ServerConfig) {
	// For dev only - Set up CORS so our client (React app) can consume the api.
	corsWrapper := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "Authorization", "Cookie", "token", "*"},
	})

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Port),
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