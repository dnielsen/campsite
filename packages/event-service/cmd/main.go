package main

import (
	"dave-web-app/packages/event-service/internal/config"
	"dave-web-app/packages/event-service/internal/database"
	"dave-web-app/packages/event-service/internal/handler"
	"dave-web-app/packages/event-service/internal/server"
	"dave-web-app/packages/event-service/internal/service"
	"dave-web-app/packages/event-service/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	// Initialize the config. If CONFIG_FILENAME isn't specified (empty string)
	// then it's gonna load the variables from environment.
	c := config.NewConfig(os.Getenv("CONFIG_FILENAME"))

	// Create a new database connection. Also, since it's a dev db,
	// migrate it and create sample mock data there.
	db := database.NewDevDb(&c.Db)

	// Create a new S3 session.
	//
	// Temporarily, we're not passing the AWS config, because there's
	// no support for it currently in our `config` module. We'll change that later.
	s := storage.NewS3Session(&c.S3)

	// We define our own HttpClient to enable mocking (for easier testing).
	var client service.HttpClient
	client = http.DefaultClient
	// Set up the API.
	api := service.NewAPI(db, s, client, c)

	// Set up the router.
	r := mux.NewRouter()

	// Register our handlers.
	r.HandleFunc("/upload", handler.UploadImage(api)).Methods(http.MethodPost)

	r.HandleFunc("/events", handler.GetAllEvents(api)).Methods(http.MethodGet)
	r.HandleFunc("/events", handler.CreateEvent(api)).Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", handler.DeleteEventById(api)).Methods(http.MethodDelete)

	r.HandleFunc("/speakers", handler.GetAllSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	r.HandleFunc("/sessions", handler.GetAllSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// Start the server.
	server.Start(&c.Server, r)
}
