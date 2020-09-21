package main

import (
	"dave-web-app/packages/session-service/internal/config"
	"dave-web-app/packages/session-service/internal/handler"
	"dave-web-app/packages/session-service/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
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
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	//// Migrate the database.
	//if err = db.AutoMigrate(&service.Session{}); err != nil {
	//	log.Fatalf("Failed to auto migrate: %v", err)
	//}
	// For dev - create mock data in the database.
	// ---------
	//now := time.Now()
	//later := now.Add(time.Hour * 1)
	//session := service.Session{
	//	ID:          "71742331-8f81-40a1-a3a1-b4c2e70160f4",
	//	Name:        "Session Name",
	//	StartDate:   &now,
	//	EndDate:     &later,
	//	Description: "description of the session",
	//	SpeakerIds: []string{"bf432767-0830-4b84-a9d2-651f2b3e7ac8"},
	//	EventId: "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
	//	Url: "https://www.youtube.com/watch?v=tTHKyJUqP44",
	//}
	//if err := db.Create(&session).Error; err != nil {
	//	log.Printf("Failed to create mock session: %v", err)
	//} else {
	//	log.Println("Created mock session in database")
	//}
	// ----------

	// We use our custom HttpClient to enable mocking.
	var client service.HttpClient
	client = http.DefaultClient
	// Set up the API.
	api := service.NewAPI(db, client, c)

	// Set up the router.
	r := mux.NewRouter()

	// Set up the handlers.
	r.HandleFunc("/", handler.GetAllSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Service.Session.Port),
		Handler:      r,
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
