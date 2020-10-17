package main

import (
	"campsite/services/session/internal/config"
	"campsite/services/session/internal/database"
	"campsite/services/session/internal/handler"
	"campsite/services/session/internal/middleware"
	"campsite/services/session/internal/server"
	"campsite/services/session/internal/service"
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
	api := service.NewAPI(db, c)

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
	r.HandleFunc("/{id}/comments", handler.CreateComment(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}/comments", handler.GetCommentsBySessionId(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.GetAllSessions(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSession(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSessionById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.EditSessionById(api)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handler.DeleteSessionById(api)).Methods(http.MethodDelete)

	// Set up and start the server.
	server.Start(r, &c.Server)
}
