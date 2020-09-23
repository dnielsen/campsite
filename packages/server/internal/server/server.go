package server

import (
	"dave-web-app/packages/server/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

const (
	READ_TIMEOUT = 15 * time.Second
	WRITE_TIMEOUT = 15 * time.Second
	IDLE_TIMEOUT = 120 * time.Second
)

func Start(c *config.ServerConfig, r *mux.Router) {
	// For dev only - Set up CORS so our client (React app) can consume the api.
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
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
