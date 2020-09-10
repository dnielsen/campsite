package main

import (
	"fmt"
	"github.com/dnielsen/campsite/packages/session-service@latest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

const PORT = 8080


type Config struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SSLMode  string `json:"sslmode"`
}

func main() {
	db, err := gorm.Open(postgres.Open(connString))
	mux := http.DefaultServeMux

	// Set up handlers.
	mux.Handle("/speakers", handler.GetSpeakers())
	// The handler serving our static files.

	// Set up the server.
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", PORT),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server.
	log.Printf("Listening on port %v", PORT)
	err := server.ListenAndServe()
	log.Fatalf("Failed to listen: %v", err)
}