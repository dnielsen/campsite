package main

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/database"
	"github.com/dnielsen/campsite/pkg/middleware"
	"github.com/dnielsen/campsite/services/api/handler"
	"github.com/dnielsen/campsite/services/api/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

const (
	SERVICE_NAME = "api"
	READ_TIMEOUT  = 15 * time.Second
	WRITE_TIMEOUT = 15 * time.Second
	IDLE_TIMEOUT  = 120 * time.Second
)

func main() {
	// Initialize the config which includes
	// Server, and other services' configuration.
	c := config.NewConfig()

	// We're running `database.NewDevDb` here so that `GORM` migrates the database for us
	// and creates mock events there.
	// It seems the least confusing to put it here rather than say the event or speaker service.
	_ = database.NewDevDb(&c.Db)

	// Set up the API.
	api := service.NewAPI(c)

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
	if c.Tracing.Enabled == true {
		r.Use(middleware.Tracing(SERVICE_NAME, c.Service.API.Port, c))
		log.Println("Tracing middleware has been enabled")
	}

	// Register our handlers.

	// UploadImage handler reads the form data and saves the retrieved image
	// into `images` directory placed in the `event` directory.
	r.HandleFunc("/images", handler.UploadImage(api)).Methods(http.MethodPost)
	// GetImage handler retrieves the image from the `images` directory placed in the project root directory.
	r.HandleFunc("/images/{filename}", handler.GetImage(api)).Methods(http.MethodGet)

	r.HandleFunc("/auth/sign-in", handler.SignIn(api)).Methods(http.MethodPost)

	r.HandleFunc("/events", handler.GetAllEvents(api)).Methods(http.MethodGet)
	r.HandleFunc("/events", handler.CreateEvent(api)).Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", handler.GetEventById(api)).Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", handler.EditEventById(api)).Methods(http.MethodPut)
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
	// GetSpeakerById handler sends a `/{id}` GET request to the speaker service
	// which retrieves the speaker from the database (if exists), and sends it back
	// to the event service which then sends it to the client (browser). It returns
	// all the properties of the speaker along with sessions.
	r.HandleFunc("/speakers/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	// EditSpeakerById handler sends a `/{id}` PUT request with input body
	// to the speaker service which edits the speaker (if exists)
	// in the database. It returns the status 204 No Content and no body.
	r.HandleFunc("/speakers/{id}", handler.EditSpeakerById(api)).Methods(http.MethodPut)
	// DeleteSpeakerById handler sends a `/{id}` DELETE request the id
	// parameter to the speaker service which deletes the speaker (if exists)
	// in the database. It returns the status 204 No Content and no body.
	r.HandleFunc("/speakers/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	// For session handlers explanation look up the speaker handlers' comments.
	// They're analogical.
	r.HandleFunc("/sessions", handler.GetAllSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/sessions/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/sessions/{id}", handler.EditSessionById(api)).Methods(http.MethodPut)
	r.HandleFunc("/sessions/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// Set up the server.
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Service.API.Port),
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
