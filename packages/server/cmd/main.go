package main

import (
	"dave-web-app/packages/server/internal/config"
	"dave-web-app/packages/server/internal/handler"
	"dave-web-app/packages/server/internal/service"
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
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Connected to database")
	// Migrate the database.
	if err = db.AutoMigrate(&service.Event{}, &service.Speaker{}, &service.Session{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Auto migrated database")
	// For dev
	// ---------
	now := time.Now()
	event := service.Event{
		ID:            "e3a27b7d-b37d-4cd2-b8bd-e5bd5551077c",
		Name:          "New Event",
		Description:   "Description",
		StartDate:     &now,
		EndDate:       &now,
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
		StartDate:   &now,
		EndDate:     &now,
		Description: "desc",
		Url:         "https://google.com",
		EventID:     event.ID,
		Speakers: []service.Speaker{speaker},
	}
	session.Speakers = []service.Speaker{speaker}
	session.Event = &event
	event.Sessions = []service.Session{session}
	event.Speakers = []service.Speaker{speaker}

	//res := db.Create(&event)
	//log.Printf("Result: %v", res.Error)
	//log.Println("Created mock data in database")
	// ----------

	api := service.NewAPI(db)

	// Set up the router.
	r := mux.NewRouter()

	// Set up handlers.
	r.HandleFunc("/events", handler.GetEvents(api)).Methods(http.MethodGet)
	r.HandleFunc("/events", handler.CreateEvent(api)).Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", handler.EditEvent(api)).Methods(http.MethodPut)

	r.HandleFunc("/speakers", handler.GetSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers/{id}", handler.EditSpeaker(api)).Methods(http.MethodPut)

	r.HandleFunc("/sessions", handler.GetSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}", handler.EditSession(api)).Methods(http.MethodPut)

	// For dev only - Set up CORS so our client can consume the api
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT"},
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
	err = srv.ListenAndServe()
	log.Fatalf("Failed to listen: %v", err)
}