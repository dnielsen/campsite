package main

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/database"
	"github.com/dnielsen/campsite/pkg/middleware"
	"github.com/dnielsen/campsite/services/speaker/handler"
	"github.com/dnielsen/campsite/services/speaker/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	SERVICE_NAME = "speaker"
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
	api := service.NewAPI(db)

	// Set up the router.
	r := mux.NewRouter()

	// Logger middleware logs the incoming requests.
	// Example output: `status=200 method=GET path=/events duration=3.714408ms`
	r.Use(middleware.Logger)

	// Enable tracing - forward our requests to the zipkin server.
	if c.Tracing.Enabled == true {
		r.Use(middleware.Tracing(SERVICE_NAME, c.Service.Speaker.Port,c))
		log.Println("Tracing middleware has been enabled")
	}

	// Set up the handlers.
	r.HandleFunc("/", handler.GetAllSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.EditSpeakerById(api)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	// Set up the server.
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", c.Service.Speaker.Port),
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
