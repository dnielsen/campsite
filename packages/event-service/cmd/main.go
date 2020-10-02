package main

import (
	"campsite/packages/event-service/internal/config"
	"campsite/packages/event-service/internal/database"
	"campsite/packages/event-service/internal/handler"
	"campsite/packages/event-service/internal/middleware"
	"campsite/packages/event-service/internal/server"
	"campsite/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	// Initialize the config.
	c := config.NewConfig()

	// Create a new database connection. Also, since it's a dev db,
	// migrate it and create sample mock data there.
	db := database.NewDevDb(&c.Db)

	// Set up the API.
	api := service.NewAPI(db, c)

	// Set up the router.
	r := mux.NewRouter()

	// Logger middleware logs the incoming requests.
	// Example output: `status=200 method=GET path=/events duration=3.714408ms`
	r.Use(middleware.Logger)
	// Request limiter middleware ensures that a client
	// with a given IP can only make so many requests.
	// If he does too many within a certain amount of time,
	// the server responds with the status code 429 (Too Many Requests).
	// This middleware is only used in the Event Service,
	// since this service is the only service that can call
	// all the others.
	r.Use(middleware.RequestLimiter)

	// Enable tracing - forward our requests to the zipkin server.
	if c.Server.Tracing.Enabled == true {
		r.Use(middleware.Tracing(&c.Server))
		log.Println("Tracing middleware has been enabled")
	}

	// Register our handlers.
	r.HandleFunc("/images", handler.UploadImage(api)).Methods(http.MethodPost)
	r.HandleFunc("/images/{filename}", handler.GetImage(api)).Methods(http.MethodGet)

	r.HandleFunc("/events", handler.GetAllEvents(api)).Methods(http.MethodGet)
	r.HandleFunc("/events", handler.CreateEvent(api)).Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", handler.EditEventById(api)).Methods(http.MethodPut)
	r.HandleFunc("/events/{id}", handler.DeleteEventById(api)).Methods(http.MethodDelete)

	r.HandleFunc("/speakers", handler.GetAllSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers/{id}", handler.EditSpeakerById(api)).Methods(http.MethodPut)
	r.HandleFunc("/speakers/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	r.HandleFunc("/sessions", handler.GetAllSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}", handler.EditSessionById(api)).Methods(http.MethodPut)
	r.HandleFunc("/sessions/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// Start the server.
	server.Start(r, &c.Server)
}
