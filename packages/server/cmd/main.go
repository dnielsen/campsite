package main

import (
	"dave-web-app/packages/server/internal/config"
	"dave-web-app/packages/server/internal/handler"
	"dave-web-app/packages/server/internal/service"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// TODO: move it to env/config
	AWS_S3_REGION = "eu-central-1"
)

func main() {
	// Initialize the config. If CONFIG_FILENAME isn't specified (empty string)
	// then it's gonna load the variables from environment.
	configFilename := os.Getenv("CONFIG_FILENAME")
	c, err := config.GetConfig(configFilename)
	if err != nil {
		log.Fatalf("Failed to load config: %v",err)
	} else {
		log.Printf("Config has been loaded: %v", c)
	}

	// Connect to the database.
	connStr := config.GetDbConnString(&c.Db)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	} else {
		log.Println("Connected to database")
	}

	// Migrate the database.
	if err = db.AutoMigrate(&service.Event{}, &service.Speaker{}, &service.Session{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	} else {
		log.Println("Auto migrated database")
	}
	// For dev - create a mock event in the database.
	// ---------
	now := time.Now()
	later := time.Now().Add(time.Hour * 8)
	address := "San Francisco, California"
	event := service.Event{
		ID:            "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Name:          "Great Event",
		Description: "Very interesting",
		StartDate:     &now,
		EndDate:       &later,
		Photo:         "https://images.unsplash.com/photo-1519834785169-98be25ec3f84?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80",
		OrganizerName: "John Tim",
		Address:       &address,
		Sessions: []service.Session{{
			ID:          "be13940b-c7ba-4f97-bdab-b4a47b11ffed",
			Name:        "Session",
			StartDate:   &now,
			EndDate:     &later,
			Description: "desc",
			Url:         "url",
			EventID:     "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
			Speakers: []service.Speaker{{
				ID:       "9c08fbf8-160b-4a86-9981-aeddf4e3798e",
				Name:     "John Doe",
				Bio:      "Bio",
				Headline: "Headline",
				Photo:    "photo",
			}},
		}},
	}

	if err := db.Create(&event).Error; err != nil {
		log.Printf("Failed to create mock data in database: %v", err)
	} else {
		log.Println("Created mock data in database")
	}
	// ----------

	s, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_S3_REGION),
		})
	if err != nil {
		log.Fatalf("Failed to create new aws session: %v", err)
	}

	// Set up the API.
	api := service.NewAPI(db, s)

	// Set up the router.
	r := mux.NewRouter()

	r.Handle("/images", handler.UploadImage(api)).Methods(http.MethodPost)
	//r.Handle("/images/{filename}", handler.GetImage()).Methods(http.MethodGet)

	// Set up the handlers.
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
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Server.Port),
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
