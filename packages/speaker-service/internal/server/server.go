package server

import (
	"dave-web-app/packages/speaker-service/internal/config"
	"dave-web-app/packages/speaker-service/internal/tracing"
	"fmt"
	"github.com/gorilla/mux"
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
	// Enable tracing, that is add a tracing middleware
	// to the router.
	if c.Tracing == true {
		tracing.EnableTracing(r, c.Port)
	}

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Port),
		Handler:      r,
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