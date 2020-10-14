package main

import (
	"campsite/packages/auth/internal/config"
	"campsite/packages/auth/internal/database"
	"campsite/packages/auth/internal/handler"
	"campsite/packages/auth/internal/middleware"
	"campsite/packages/auth/internal/server"
	"campsite/packages/auth/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	// Initialize the config which includes
	// database, server, and other services' configuration.
	c := config.NewConfig()

	// Create a new database connection. Also, since it's a dev db,
	// migrate it and create sample mock data there.
	db := database.NewDb(&c.Db)

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
	// Plain JWT sign in
	r.HandleFunc("/sign-in", handler.SignIn(api)).Methods(http.MethodPost)
	// GitHub Sign In
	r.HandleFunc("/sign-in/github", handler.GitHubSignIn())
	r.HandleFunc("/sign-in/github/callback", handler.GitHubSignInCallback(api))


	// Start the server. It sets up CORS for us
	// so that our `ui` or any other client can consume the API
	// conveniently. Also it configures read, idle, and write timeouts.
	server.Start(r, &c.Server)
}