package main

import (
	"dave-web-app/packages/event-service/internal/handler"
	"dave-web-app/packages/event-service/internal/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

const PORT = 4444

func main() {
	connStr := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	if err = db.AutoMigrate(&service.Event{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	// For dev
	// ---------
	event := service.Event{
		ID:            uuid.New().String(),
		Name:          "Great Event",
		Description: "Very interesting",
		StartDate:     time.Now(),
		EndDate:       time.Date(2022, time.November, 10, 23, 0, 0, 0, time.UTC),
		Photo:         "https://images.unsplash.com/photo-1519834785169-98be25ec3f84?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80",
		OrganizerName: "John Tim",
		Address:       "San Francisco, California",
		SpeakerIds: []string{"7fd20c48-2575-4a88-91fa-34665f76c6f0"},
	}
	res := db.Create(&event)
	log.Println(event)
	log.Println(res.RowsAffected)
	// ----------

	// We use our custom HttpClient to enable mocking.
	var c service.HttpClient
	c = http.DefaultClient
	api := service.NewAPI(db, c)

	// Set up handlers.
	r := mux.NewRouter()

	r.HandleFunc("/sessions/{sessionId}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.GetAllSessions(api)).Methods(http.MethodGet)

	r.HandleFunc("/speakers/{speakerId}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.GetAllSpeakers(api)).Methods(http.MethodGet)

	r.HandleFunc("/{eventId}", handler.GetEventById(api)).Methods(http.MethodGet)
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
