package main

import (
	"../internal/handler"
	"../internal/service"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

const PORT = 8080

//type Config struct {
//	Name     string `json:"name"`
//	User     string `json:"user"`
//	Password string `json:"password"`
//	Host     string `json:"host"`
//	Port     string `json:"port"`
//	SSLMode  string `json:"sslmode"`
//}


func main() {
	connStr := "user=postgres password=postgres dbname=campsite port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	api := service.NewAPI(db)

	// Set up handlers.
	mux := http.DefaultServeMux
	mux.Handle("/speakers", handler.GetSpeakers(api))
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
	err = server.ListenAndServe()
	log.Fatalf("Failed to listen: %v", err)
}