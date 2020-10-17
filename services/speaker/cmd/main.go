package main

import (
	"campsite/packages/speaker/internal/config"
	"campsite/packages/speaker/internal/database"
	"campsite/packages/speaker/internal/handler"
	"campsite/packages/speaker/internal/middleware"
	"campsite/packages/speaker/internal/server"
	"campsite/packages/speaker/internal/service"
	"github.com/gorilla/mux"
	"log"
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

	// Logger middleware logs the incoming requests.
	// Example output: `status=200 method=GET path=/events duration=3.714408ms`
	r.Use(middleware.Logger)

	// Enable tracing - forward our requests to the zipkin server.
	if c.Server.Tracing.Enabled == true {
		r.Use(middleware.Tracing(&c.Server))
		log.Println("Tracing middleware has been enabled")
	}

	// Set up the handlers.
	r.HandleFunc("/", handler.GetAllSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.EditSpeakerById(api)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	// Set up and start the server.
	server.Start(r, &c.Server)
}
