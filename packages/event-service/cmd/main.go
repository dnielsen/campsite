package main

import (
	"dave-web-app/packages/event-service/internal/config"
	"dave-web-app/packages/event-service/internal/handler"
	"dave-web-app/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	// Initialize the config. If it can't find the file, it will load the variables
	// from the environment. It would be a good idea to read the file path to the config
	// from environment, because we might want to have `test.yml` or some other config.
	c, err := config.GetConfig("development.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v",err)
	}
	log.Printf("Config has been loaded: %v", c)

	// Connect to the database.
	connStr := config.GetDbConnString(&c.Db)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	// Migrate the database.
	if err = db.AutoMigrate(&service.Event{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	// For dev
	// ---------
	event := service.Event{
		ID:            "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Name:          "Great Event",
		Description: "Very interesting",
		StartDate:     time.Now(),
		EndDate:       time.Date(2022, time.November, 10, 23, 0, 0, 0, time.UTC),
		Photo:         "https://images.unsplash.com/photo-1519834785169-98be25ec3f84?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80",
		OrganizerName: "John Tim",
		Address:       "San Francisco, California",
		SpeakerIds: []string{"bf432767-0830-4b84-a9d2-651f2b3e7ac8"},
	}
	res := db.Create(&event)
	log.Println(event)
	log.Println(res.RowsAffected)
	// ----------

	// We use our custom HttpClient to enable mocking.
	var client service.HttpClient
	client = http.DefaultClient
	api := service.NewAPI(db, client)

	// Set up handlers.
	r := mux.NewRouter()

	r.HandleFunc("/sessions/{sessionId}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.GetAllSessions(api)).Methods(http.MethodGet)

	r.HandleFunc("/speakers/{speakerId}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.GetAllSpeakers(api)).Methods(http.MethodGet)

	r.HandleFunc("/{eventId}", handler.GetEventById(api)).Methods(http.MethodGet)
	// Set up the server.
	srv := &http.Server{
		Addr:         c.Server.Address,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server.
	log.Printf("Listening at: %v", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatalf("Failed to listen: %v", err)
}
