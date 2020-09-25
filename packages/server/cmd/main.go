package main

import (
	"dave-web-app/packages/server/internal/config"
	"dave-web-app/packages/server/internal/database"
	"dave-web-app/packages/server/internal/handler"
	"dave-web-app/packages/server/internal/server"
	"dave-web-app/packages/server/internal/service"
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

	// Set up the API.
	api := service.NewAPI(db, c)

	// Set up the router.
	r := mux.NewRouter()

	// Register our handlers.
	r.HandleFunc("/images", handler.UploadImage(api)).Methods(http.MethodPost)
	r.HandleFunc("/images/{filename}", handler.GetImage(api)).Methods(http.MethodGet)

	r.HandleFunc("/events", handler.GetEvents(api)).Methods(http.MethodGet)
	r.HandleFunc("/events", handler.CreateEvent(api)).Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", handler.DeleteEventById(api)).Methods(http.MethodDelete)
	r.HandleFunc("/events/{id}", handler.EditEventById(api)).Methods(http.MethodPut)

	r.HandleFunc("/speakers", handler.GetSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)
	r.HandleFunc("/sessions/{id}", handler.EditSpeakerById(api)).Methods(http.MethodPut)

	r.HandleFunc("/sessions", handler.GetSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)
	r.HandleFunc("/sessions/{id}", handler.EditSessionById(api)).Methods(http.MethodPut)

	// Start the server.
	server.Start(&c.Server, r)
}
