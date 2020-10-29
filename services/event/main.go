package main

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/database"
	"github.com/dnielsen/campsite/pkg/middleware"
	"github.com/dnielsen/campsite/pkg/tracing"
	"github.com/dnielsen/campsite/services/event/handler"
	"github.com/dnielsen/campsite/services/event/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	SERVICE_NAME = "event"
	READ_TIMEOUT  = 15 * time.Second
	WRITE_TIMEOUT = 15 * time.Second
	IDLE_TIMEOUT  = 120 * time.Second
)

func main() {
	// Initialize the config which includes
	// database, server, and other services' configuration.
	c := config.NewConfig()

	// Create a new database connection.
	db := database.NewDevDb(c)

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

	// Enable tracing middleware - forward our request data to the zipkin server
	// that is running with Hypertrace.
	// Enable tracing - forward our requests to the zipkin server.
	if c.Tracing.Enabled == true {
		tracer := tracing.NewTracer(SERVICE_NAME, strconv.Itoa(c.Service.Event.Port), c)
		r.Use(middleware.Tracing(tracer))
		log.Println("Tracing has been enabled")
	}
	// Register our handlers.
	// GetAllEvents handler selects all events along with all the properties
	// from the database and sends them to the client. It doesn't join any tables.
	// We could optimize this so that it would skip the `EndDate` property
	// since our `ui` isn't using it currently.
	r.HandleFunc("/", handler.GetAllEvents(api)).Methods(http.MethodGet)
	// CreateEvent handler creates an event in the database.
	// If the event creation succeeds, it sends the created event and status 201 (Created).
	// It doesn't validate the input currently.
	r.HandleFunc("/", handler.CreateEvent(api)).Methods(http.MethodPost)

	// GetEventById handler retrieves an event with a given id from the database.
	// If it can't find it, it's gonna return an error, and send status 404 (Not Found).
	// We could optimize this by just returning the id of the created event
	// since our `ui` isn't using this data besides the `id` to redirect
	// to the created event.
	r.HandleFunc("/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	// EditEventById handler edits an event with a given id in the database.
	// It sends a status 204 (No Content) if the edit has been performed successfully.
	// It doesn't return the updated event. If the event couldn't be found, it's gonna
	// return an error, and status 404 (Not Found).
	// It doesn't validate the input currently.
	r.HandleFunc("/{id}", handler.EditEventById(api)).Methods(http.MethodPut)



	// DeleteEventById handler deletes an event with a given id in the database.
	// It sends a status 204 (No Content) if the delete has been performed successfully.
	// It doesn't return the deleted event. If the event couldn't be found, it's gonna
	// return an error, and status 404 (Not Found).
	r.HandleFunc("/{id}", handler.DeleteEventById(api)).Methods(http.MethodDelete)




	// Set up the server.
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
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
