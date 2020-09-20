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
		log.Fatalf("Failed to load config: %v", err)
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
		ID:            "aef5329b-b934-4d60-bf33-ff2ec368c119",
		Name:          "BigDataCamp LA 2020",
		Description:   "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor, excepturi magnam nisi numquam quam quidem soluta voluptatem. Aspernatur aut, consequuntur dolore et laudantium libero magnam officia quod repellendus ullam, voluptatibus.",
		StartDate:     &now,
		EndDate:       &now,
		Photo:         "https://devconf.info/assets/images/devconf-cz-social.png",
		OrganizerName: "David Musk",
		Address:       "San Jose, CA",
	}
	speaker := service.Speaker{
		ID:       "aef5329b-b934-4d60-bf33-ff2ec368c119",
		Name:     "John Doe",
		Bio:      "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor, excepturi magnam nisi numquam quam quidem soluta voluptatem. Aspernatur aut, consequuntur dolore et laudantium libero magnam officia quod repellendus ullam, voluptatibus.",
		Headline: "CEO of Tesla",
		Photo:    "https://www.mantruckandbus.com/fileadmin/_processed_/2/c/csm_kleiss-interview-header_8a76bdbcb6.jpg",
	}
	session := service.Session{
		ID:          "391700d5-08dc-4173-9193-80ea1a32b7f9",
		Name:        "Concurrency in Go",
		StartDate:   &now,
		EndDate:     &now,
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor, excepturi magnam nisi numquam quam quidem soluta voluptatem. Aspernatur aut, consequuntur dolore et laudantium libero magnam officia quod repellendus ullam, voluptatibus.",
		Url:         "https://google.com",
		EventID:     &event.ID,
		Speakers:    []service.Speaker{speaker},
	}
	session.Speakers = []service.Speaker{speaker}
	event.Sessions = []service.Session{session}

	if err := db.Create(&event).Error; err != nil {
		log.Printf("Failed to create mock data in database: %v", err)
	} else {
		log.Println("Created mock data in database")
	}
	// ----------

	api := service.NewAPI(db)

	// Set up the router.
	r := mux.NewRouter()

	// Set up handlers.
	r.HandleFunc("/events", handler.GetEvents(api)).Methods(http.MethodGet)
	r.HandleFunc("/events", handler.CreateEvent(api)).Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", handler.DeleteEvent(api)).Methods(http.MethodDelete)

	r.HandleFunc("/speakers", handler.GetSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers/{id}", handler.DeleteSpeaker(api)).Methods(http.MethodDelete)

	r.HandleFunc("/sessions", handler.GetSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}", handler.DeleteSession(api)).Methods(http.MethodDelete)

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
	err = srv.ListenAndServe()
	log.Fatalf("Failed to listen: %v", err)
}
