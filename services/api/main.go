package main

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/middleware"
	"github.com/dnielsen/campsite/pkg/tracing"
	"github.com/dnielsen/campsite/services/api/handler"
	"github.com/gorilla/mux"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
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
	client := http.DefaultClient

	// We're running `database.NewDevDb` here so that `GORM` migrates the database for us
	// and creates mock events there.
	// It seems the least confusing to put it here rather than say the event or speaker service.
	//_ = database.NewDevDb(c)

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
	if c.Tracing.Enabled {
		tracer := tracing.NewTracer(SERVICE_NAME, strconv.Itoa(c.Service.API.Port), c)

		r.Use(middleware.Tracing(tracer))
		transport, err := zipkinHttp.NewTransport(tracer, zipkinHttp.TransportTrace(true))
		if err != nil {
			log.Fatalf("Failed to create zipkin transport")
		}
		client = &http.Client{Transport: transport}
		log.Println("Tracing has been enabled")
	}

	// Register our handlers.

	// UploadImage handler reads the form data and saves the retrieved image
	// into `images` directory placed in the `event` directory.
	r.HandleFunc("/api/images", handler.UploadImage(c)).Methods(http.MethodPost)
	// GetImageByFilename handler retrieves the image from the `images` directory placed in the project root directory.
	r.HandleFunc("/api/images/{filename}", handler.GetImageByFilename(c)).Methods(http.MethodGet)

	// If the user is signed in (has the access token) it returns a `Me` struct
	// with the user data such as `ID`, `Email`. Otherwise it returns an empty response.
	// Either way the status code should be 200.
	r.HandleFunc("/api/auth", handler.Auth(c)).Methods(http.MethodGet)
	r.HandleFunc("/api/auth/sign-in", handler.SignIn(client, c)).Methods(http.MethodPost)
	r.HandleFunc("/api/auth/sign-out", handler.SignOut(c)).Methods(http.MethodPost)

	r.HandleFunc("/api/events", handler.GetAllEvents(client, c)).Methods(http.MethodGet)
	r.HandleFunc("/api/events", handler.CreateEvent(client, c)).Methods(http.MethodPost)
	r.HandleFunc("/api/events/{id}", handler.GetEventById(client, c)).Methods(http.MethodGet)
	r.HandleFunc("/api/events/{id}", handler.EditEventById(client, c)).Methods(http.MethodPut)
	r.HandleFunc("/api/events/{id}", handler.DeleteEventById(client, c)).Methods(http.MethodDelete)

	r.HandleFunc("/api/speakers", handler.GetAllSpeakers(client, c)).Methods(http.MethodGet)
	r.HandleFunc("/api/speakers", handler.CreateSpeaker(client, c)).Methods(http.MethodPost)
	r.HandleFunc("/api/speakers/{id}", handler.GetSpeakerById(client, c)).Methods(http.MethodGet)
	r.HandleFunc("/api/speakers/{id}", handler.EditSpeakerById(client, c)).Methods(http.MethodPut)
	r.HandleFunc("/api/speakers/{id}", handler.DeleteSpeakerById(client, c)).Methods(http.MethodDelete)

	r.HandleFunc("/api/sessions", handler.GetAllSessions(client, c)).Methods(http.MethodGet)
	r.HandleFunc("/api/sessions", handler.CreateSession(client, c)).Methods(http.MethodPost)
	r.HandleFunc("/api/sessions/{id}", handler.GetSessionById(client, c)).Methods(http.MethodGet)
	r.HandleFunc("/api/sessions/{id}", handler.EditSessionById(client, c)).Methods(http.MethodPut)
	r.HandleFunc("/api/sessions/{id}", handler.DeleteSessionById(client, c)).Methods(http.MethodDelete)

	// Set up the server.
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:         []string{"http://localhost:3000"},
		AllowedMethods:         []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:         []string{"Content-Type", "Origin", "Accept", "Authorization", "Cookie", "Set-Cookie"},
		AllowCredentials:       true,
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
