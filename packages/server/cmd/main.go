package main

import (
	"dave-web-app/packages/server/internal/config"
	"dave-web-app/packages/server/internal/handler"
	"dave-web-app/packages/server/internal/service"
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
	if err = db.AutoMigrate(&service.Event{}, &service.Speaker{}, &service.Session{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	// For dev
	// ---------
	event := service.Event{
		ID:            "e3a27b7d-b37d-4cd2-b8bd-e5bd5551077c",
		Name:          "New Event",
		Description:   "Description",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		Photo:         "https://google.com",
		OrganizerName: "David Musk",
		Address:       "San Jose, CA",
	}
	speaker := service.Speaker{
		ID:       "aef5329b-b934-4d60-bf33-ff2ec368c119",
		Name:     "Speaker name",
		Bio:      "Speaker bio",
		Headline: "Speaker headline",
		Photo:    "https://youtube.com",
	}
	session := service.Session{
		ID:          "391700d5-08dc-4173-9193-80ea1a32b7f9",
		Name:        "session name",
		StartDate:   time.Now(),
		EndDate:     time.Now(),
		Description: "desc",
		Url:         "https://google.com",
		EventID:     event.ID,
		Speakers: []service.Speaker{speaker},
	}
	session.Speakers = []service.Speaker{speaker}
	event.Sessions = []service.Session{session}
	event.Speakers = []service.Speaker{speaker}

	_ = db.Create(&event)
	// ----------

	api := service.NewAPI(db)

	// Set up handlers.
	r := mux.NewRouter()

	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
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