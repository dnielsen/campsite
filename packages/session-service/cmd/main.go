package main

import (
	"dave-web-app/packages/session-service/internal/handler"
	"dave-web-app/packages/session-service/internal/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

const PORT = 5555

func main() {
	connStr := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	if err = db.AutoMigrate(&service.Session{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	// For dev
	// ---------
	session := service.Session{
		ID:          uuid.New().String(),
		Title:       "Session Title",
		StartDate:   time.Now(),
		EndDate:     time.Now().AddDate(1, 1, 1),
		Description: "description of the session",
	}
	res := db.Create(&session)
	log.Println(session)
	log.Println(res.RowsAffected)
	// ----------

	api := service.NewAPI(db)

	// Set up handlers.
	r := mux.NewRouter()
	r.HandleFunc("/events/{id}", handler.GetSessionsByIds(api)).Methods(http.MethodGet)

	// Set up the server.
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", PORT),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server.
	log.Printf("Listening on port %v", PORT)
	err = server.ListenAndServe()
	log.Fatalf("Failed to listen: %v", err)
}
