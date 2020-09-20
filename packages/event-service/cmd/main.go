package main

import (
	"dave-web-app/packages/event-service/internal/config"
	"dave-web-app/packages/event-service/internal/handler"
	"dave-web-app/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	now := time.Now()
	later := time.Now().Add(time.Hour * 8)
	event := service.Event{
		ID:            "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Name:          "Great Event",
		Description: "Very interesting",
		StartDate:     &now,
		EndDate:       &later,
		Photo:         "https://images.unsplash.com/photo-1519834785169-98be25ec3f84?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80",
		OrganizerName: "John Tim",
		Address:       "San Francisco, California",
	}
	if err := db.Create(&event).Error; err != nil {
		log.Printf("Failed to create mock event in database: %v", err)
	} else {
		log.Println("Created mock event in database")
	}
	// ----------

	// We use our custom HttpClient to enable mocking.
	var client service.HttpClient
	client = http.DefaultClient
	api := service.NewAPI(db, client, c)

	// Set up handlers.
	r := mux.NewRouter()

	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.GetAllSessions(api)).Methods(http.MethodGet)

	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.GetAllSpeakers(api)).Methods(http.MethodGet)

	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	r.HandleFunc("/events", handler.GetEvents(api)).Methods(http.MethodGet)

	// For dev only - Set up CORS so our client can consume the api
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	// Set up the server.
	srv := &http.Server{
		Addr:         c.Server.Address,
		Handler:      corsWrapper.Handler(r),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server.
	log.Printf("Listening at: %v", srv.Addr)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
