package database

import (
	"campsite/services/speaker/internal/config"
	"campsite/services/speaker/internal/util"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
)

// Example output:
// `"user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"`
func getDbConnString(c *config.DbConfig) string {
	vals := getDbValues(c)
	var p []string
	for k, v := range vals {
		p = append(p, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(p, " ")
}

func getDbValues(c *config.DbConfig) map[string]string {
	p := map[string]string{}
	util.SetIfNotEmpty(p, "dbname", c.Name)
	util.SetIfNotEmpty(p, "host", c.Host)
	util.SetIfNotEmpty(p, "user", c.User)
	util.SetIfNotEmpty(p, "password", c.Password)
	util.SetIfNotEmpty(p, "port", c.Port)
	util.SetIfNotEmpty(p, "sslmode", c.SSLMode)
	return p
}

// Initializes a new database with the provided configuration.
// Crashes the program if failed to connect.
func NewDb(c *config.DbConfig) *gorm.DB {
	connStr := getDbConnString(c)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")
	return db
}
