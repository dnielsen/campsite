package main

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/database"
	"github.com/dnielsen/campsite/pkg/middleware"
	"github.com/dnielsen/campsite/pkg/tracing"
	"github.com/dnielsen/campsite/services/session/handler"
	"github.com/dnielsen/campsite/services/session/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const(
	SERVICE_NAME = "session"
	READ_TIMEOUT  = 15 * time.Second
	WRITE_TIMEOUT = 15 * time.Second
	IDLE_TIMEOUT  = 120 * time.Second
)

func main() {
	// Initialize the config.
	c := config.NewConfig()

	// Create a new database connection.
	db := database.NewDb(&c.Db)

	// Set up the API.
	api := service.NewAPI(db, c)

	// Set up the router.
	r := mux.NewRouter()

	// Logger middleware logs the incoming requests.
	// Example output: `status=200 method=GET path=/events duration=3.714408ms`
	r.Use(middleware.Logger)

	// Enable tracing - forward our requests to the zipkin server.
	if c.Tracing.Enabled == true {
		tracer := tracing.NewTracer(SERVICE_NAME, string(rune(c.Service.Session.Port)), c)
		r.Use(middleware.Tracing(tracer))
		log.Println("Tracing has been enabled")
	}

	// Set up the handlers.
	r.HandleFunc("/", handler.GetAllSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.EditSessionById(api)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// Set up and start the server.
	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Service.Session.Port),
		Handler:      r,
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
