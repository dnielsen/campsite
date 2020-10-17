package main

import (
	"campsite/pkg/config"
	"campsite/pkg/database"
	"campsite/pkg/middleware"
	"campsite/services/event/handler"
	"campsite/services/event/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

const (
	// Service name for tracing.
	SERVICE_NAME = "event"
	// Server config
	READ_TIMEOUT  = 15 * time.Second
	WRITE_TIMEOUT = 15 * time.Second
	IDLE_TIMEOUT  = 120 * time.Second
)

func main() {

	// Initialize the config which includes
	// database, server, and other services' configuration.
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
	if c.Tracing.Enabled == true {
		r.Use(middleware.Tracing(SERVICE_NAME, c.Service.Event.Port, c))
		log.Println("Tracing middleware has been enabled")
	}

	// Register our handlers.
	r.HandleFunc("/auth/sign-in", handler.SignIn(api)).Methods(http.MethodPost)
	// UploadImage handler reads the form data and saves retrieved image
	// into `images` directory placed in the `event` directory.
	r.HandleFunc("/images", handler.UploadImage(api)).Methods(http.MethodPost)
	// GetImage handler retrieves the image from the `images` directory
	// placed in the `event` directory.
	r.HandleFunc("/images/{filename}", handler.GetImage(api)).Methods(http.MethodGet)
	// GetAllEvents handler selects all events along with all the properties
	// from the database and sends them to the client. It doesn't join any tables.
	// We could optimize this so that it would skip the `EndDate` property
	// since our `ui` isn't using it currently.
	r.HandleFunc("/events", handler.GetAllEvents(api)).Methods(http.MethodGet)
	// CreateEvent handler creates an event in the database.
	// If the event creation succeeds, it sends the created event and status 201 (Created).
	// It doesn't validate the input currently.
	r.HandleFunc("/events", handler.CreateEvent(api)).Methods(http.MethodPost)
	// GetEventById handler retrieves an event with a given id from the database.
	// If it can't find it, it's gonna return an error, and send status 404 (Not Found).
	// We could optimize this by just returning the id of the created event
	// since our `ui` isn't using this data besides the `id` to redirect
	// to the created event.
	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	// EditEventById handler edits an event with a given id in the database.
	// It sends a status 204 (No Content) if the edit has been performed successfully.
	// It doesn't return the updated event. If the event couldn't be found, it's gonna
	// return an error, and status 404 (Not Found).
	// It doesn't validate the input currently.
	r.HandleFunc("/events/{id}", handler.EditEventById(api)).Methods(http.MethodPut)
	// DeleteEventById handler deletes an event with a given id in the database.
	// It sends a status 204 (No Content) if the delete has been performed successfully.
	// It doesn't return the deleted event. If the event couldn't be found, it's gonna
	// return an error, and status 404 (Not Found).
	r.HandleFunc("/events/{id}", handler.DeleteEventById(api)).Methods(http.MethodDelete)

	// GetAllSpeakers handler sends a `/` GET request to the speaker service
	// which selects all the speakers along with all the properties
	// from the database and sends them back to the event service which
	// then sends them to the client (browser for example). It doesn't join any tables.
	// We could optimize this so that it would skip the `bio` property since
	// it's not used by our `ui`.
	r.HandleFunc("/speakers", handler.GetAllSpeakers(api)).Methods(http.MethodGet)
	// CreateSpeaker handler sends a `/` POST request with input body
	// to the speaker service which creates a speaker in the database,
	// and sends the newly created speaker back to the event service which
	// then sends them to the client (browser for example).
	// There's currently no input validation.
	// We could optimize this by just returning the id of the created speaker
	// since our `ui` isn't using this data besides the `id` to redirect
	// to the created speaker.
	r.HandleFunc("/speakers", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/speakers/{id}", handler.EditSpeakerById(api)).Methods(http.MethodPut)
	r.HandleFunc("/speakers/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	r.HandleFunc("/sessions/{id}/comments", handler.GetCommentsBySessionId(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}/comments", handler.CreateComment(api)).Methods(http.MethodPost)
	// GetAllSessions handler selects all sessions along with all the properties
	// from the database and sends them to the client. It doesn't join any tables.
	// It's currently not being used by our `ui`.
	r.HandleFunc("/sessions", handler.GetAllSessions(api)).Methods(http.MethodGet)
	// CreateSession handler sends a `/` POST request with input body
	// to the session service which creates a session in the database,
	// and sends the newly created session back to the event service which
	// then sends them to the client (browser for example).
	// There's currently no input validation.
	// We could optimize this by just returning the id of the created session
	// since our `ui` isn't using this data besides the `id` to redirect
	// to the created session.
	r.HandleFunc("/sessions", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}", handler.EditSessionById(api)).Methods(http.MethodPut)
	r.HandleFunc("/sessions/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// For dev only - Set up CORS so our client (React app) can consume the api.
	corsWrapper := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "Authorization", "Cookie", "token", "*"},
	})

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Service.Event.Port),
		Handler:      corsWrapper.Handler(r),
		ReadTimeout:  READ_TIMEOUT,
		WriteTimeout: WRITE_TIMEOUT,
		IdleTimeout:  IDLE_TIMEOUT,
	}

	// Start the server.
	log.Printf("Listening at: %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
