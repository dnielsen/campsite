package main

import (
	"campsite/packages/speaker-service/internal/config"
	"campsite/packages/speaker-service/internal/database"
	"campsite/packages/speaker-service/internal/handler"
	"campsite/packages/speaker-service/internal/server"
	"campsite/packages/speaker-service/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Initialize the config.
	c := config.NewConfig()

	// Create a new database connection.
	db := database.NewDb(&c.Db)

	// Set up the API.
	api := service.NewAPI(db)

	// Set up the router.
	r := mux.NewRouter()

	// Set up the handlers.
	r.HandleFunc("/", handler.GetAllSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.EditSpeakerById(api)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	// Set up and start the server.
	server.Start(r, &c.Server)
}