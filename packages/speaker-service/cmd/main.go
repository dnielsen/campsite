package main

import (
	"dave-web-app/packages/speaker-service/internal/handler"
	"dave-web-app/packages/speaker-service/internal/service"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

const PORT = 3333

//type Config struct {
//	Name     string `json:"name"`
//	User     string `json:"user"`
//	Password string `json:"password"`
//	Host     string `json:"host"`
//	Port     string `json:"port"`
//	SSLMode  string `json:"sslmode"`
//}


func main() {
	connStr := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	if err = db.AutoMigrate(&service.Speaker{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	// For dev
	// ---------
	speaker := service.Speaker{Name: "John Doe"}
	result := db.Create(&speaker)
	log.Println(speaker)
	log.Println(result.RowsAffected)
	// ----------

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