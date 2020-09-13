package main

import (
	"dave-web-app/packages/speaker-service/internal/handler"
	"dave-web-app/packages/speaker-service/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

const PORT = 3333

func main() {
	// Temporary solution
	connStr := "host=host.docker.internal user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	if err = db.AutoMigrate(&service.Speaker{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	// For dev
	// ---------
	speaker := service.Speaker{
		ID: 		"bf432767-0830-4b84-a9d2-651f2b3e7ac8",
		Name:     "Warren Josh",
		Bio:      "I'm a computer geek",
		Headline: "CEO of Hello",
		Photo:    "https://images.unsplash.com/photo-1546661717-0bf190068ede?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjF9&auto=format&fit=crop&w=1525&q=80",
	}
	result := db.Create(&speaker)
	log.Println(speaker)
	log.Println(result.RowsAffected)
	// ----------

	api := service.NewAPI(db)

	// Set up handlers.
	r := mux.NewRouter()
	r.HandleFunc("/byIds", handler.GetSpeakersByIds(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.GetAllSpeakers(api)).Methods(http.MethodGet)

	// Set up the server.
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", PORT),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server.
	log.Printf("Listening on port %v", PORT)
	err = server.ListenAndServe()
	log.Fatalf("Failed to listen: %v", err)
}