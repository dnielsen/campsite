package main

import (
	"dave-web-app/packages/speaker-service/internal/config"
	"dave-web-app/packages/speaker-service/internal/handler"
	"dave-web-app/packages/speaker-service/internal/service"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	// Initialize the config. If it can't find the file, it will load the variables
	// from the environment. It would be a good idea to read the file path to the config
	// from environment, because we might want to have `test.yml` or some other config.
	c, err := config.GetConfig("development.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v",err)
	}
	log.Printf("Config has been loaded: %v", c)

	// Connect to the database.
	connStr := config.GetDbConnString(&c.Db)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	//// Migrate the database.
	//if err = db.AutoMigrate(&service.Speaker{}); err != nil {
	//	log.Fatalf("Failed to auto migrate: %v", err)
	//}
	// For dev
	// ---------
	//speaker := service.Speaker{
	//	ID: 		"bf432767-0830-4b84-a9d2-651f2b3e7ac8",
	//	Name:     "Warren Josh",
	//	Bio:      "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fuga recusandae sequi vel velit veniam? Animi assumenda dicta doloribus excepturi id ipsa natus placeat quidem ratione voluptatibus! A aut tempora veritatis.         Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fuga recusandae sequi vel velit veniam? Animi assumenda dicta doloribus excepturi id ipsa natus placeat quidem ratione voluptatibus! A aut tempora veritatis.         Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fuga recusandae sequi vel velit veniam? Animi assumenda dicta doloribus excepturi id ipsa natus placeat quidem ratione voluptatibus! A aut tempora veritatis.",
	//	Headline: "CEO of Hello",
	//	Photo:    "https://images.unsplash.com/photo-1546661717-0bf190068ede?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjF9&auto=format&fit=crop&w=1525&q=80",
	//}
	//if err := db.Create(&speaker).Error; err != nil {
	//	log.Printf("Failed to create mock speaker in database: %v", err)
	//} else {
	//	log.Println("Created mock speaker in database")
	//}
	// ----------

	// Set up the API.
	api := service.NewAPI(db)

	// Set up the router.
	r := mux.NewRouter()

	// Set up the handlers.
	r.HandleFunc("/", handler.GetSpeakers(api)).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateSpeaker(api)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.GetSpeakerById(api)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.DeleteSpeakerById(api)).Methods(http.MethodDelete)

	// Set up the server.
	srv := &http.Server{
		Addr:         c.Server.Address,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server.
	log.Printf("Listening at: %v", srv.Addr)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}