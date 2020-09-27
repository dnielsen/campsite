package main

import (
	"contrib.go.opencensus.io/exporter/zipkin"
	"dave-web-app/packages/event-service/internal/config"
	"dave-web-app/packages/event-service/internal/database"
	"dave-web-app/packages/event-service/internal/handler"
	"dave-web-app/packages/event-service/internal/server"
	"dave-web-app/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	openZipkin "github.com/openzipkin/zipkin-go"
	httpZipkin "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/trace"
	"log"
	"net/http"
)

func main() {
	localEndpointURI := "192.168.1.5:5454"
	reporterURI := "http://localhost:9411/api/v2/spans"
	serviceName := "event-service"
	localEndpoint, err := openZipkin.NewEndpoint(serviceName, localEndpointURI)
	if err != nil {
		log.Fatalf("Failed to create new zipkin endpoint: %v", err)
	}
	reporter := httpZipkin.NewReporter(reporterURI)
	exp := zipkin.NewExporter(reporter, localEndpoint)
	trace.RegisterExporter(exp)
	// Initialize the config.
	c := config.NewConfig()

	// Create a new database connection. Also, since it's a dev db,
	// migrate it and create sample mock data there.
	db := database.NewDevDb(&c.Db)

	// We define our own HttpClient to enable mocking (for easier testing).
	var client service.HttpClient
	client = http.DefaultClient
	// Set up the API.
	api := service.NewAPI(db, client, c)

	// Set up the router.
	r := mux.NewRouter()

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
	server.Start(&c.Server, r)
}
