package main

import (
	"dave-web-app/packages/session-service/internal/config"
	"dave-web-app/packages/session-service/internal/database"
	"dave-web-app/packages/session-service/internal/handler"
	"dave-web-app/packages/session-service/internal/server"
	"dave-web-app/packages/session-service/internal/service"
	"dave-web-app/packages/session-service/internal/tracing"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Initialize the config.
	c := config.NewConfig()

	// Create a new database connection.
	db := database.NewDb(&c.Db)

	// We're using our custom `HttpClient` to enable mocking.
	var client service.HttpClient
	client = http.DefaultClient
	// Set up the API.
	api := service.NewAPI(db, client, c)

	// Set up the router.
	r := mux.NewRouter()

	// Set up the tracing middleware.
	t := tracing.NewTracer(&c.Server)
	tracingMiddleware := tracing.NewTracingMiddleware(t)
	r.Use(tracingMiddleware)

	// Set up the handlers.
	r.HandleFunc("/", handler.GetAllSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.EditSessionById(api)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// Set up and start the server.
	server.Start(&c.Server, r)
}
