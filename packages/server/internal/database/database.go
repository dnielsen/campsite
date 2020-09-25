package database

import (
	"dave-web-app/packages/server/internal/config"
	"dave-web-app/packages/server/internal/service"
	"dave-web-app/packages/server/internal/util"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

func getDbConnString(c *config.DbConfig) string  {
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

func NewDb(c *config.DbConfig) *gorm.DB {
	connStr := getDbConnString(c)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")
	return db
}

// Same as NewDb but additionally migrates the database and creates
// mock data in the database.
func NewDevDb(c *config.DbConfig) *gorm.DB  {
	db := NewDb(c)
	// Migrate the database.
	if err := db.AutoMigrate(&service.Event{}, &service.Speaker{}, &service.Session{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Auto migrated database")

	// Create mock data in the database.
	event := getMockEvent()
	if err := db.Create(&event).Error; err != nil {
		// The error will likely occur because we already created it already,
		// that is, the primary keys we set up above already exists.
		// We can ignore this.
		log.Printf("Failed to create mock data in database: %v", err)
	}
	log.Println("Created mock data in database")

	return db
}

func getMockEvent() service.Event {
	now := time.Now()
	later := time.Now().Add(time.Hour * 8)
	address := "San Francisco, California"
	event := service.Event{
		ID:            "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Name:          "BigDataCamp LA 2020",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Asperiores commodi corporis cumque eaque inventore minima neque nostrum odit perspiciatis, quam ut veritatis voluptas voluptatibus. Adipisci atque deleniti fugiat odit ullam!",
		StartDate:     &now,
		EndDate:       &later,
		Photo:         "https://www.phplist.org/wp-content/uploads/2019/12/devconf2020-960x411.png",
		OrganizerName: "John Tim",
		Address:       &address,
		Sessions: []service.Session{{
			ID:          "be13940b-c7ba-4f97-bdab-b4a47b11ffed",
			Name:        "Intro to Python",
			StartDate:   &now,
			EndDate:     &later,
			Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Asperiores commodi corporis cumque eaque inventore minima neque nostrum odit perspiciatis, quam ut veritatis voluptas voluptatibus. Adipisci atque deleniti fugiat odit ullam!",
			Url:         "https://apple.com",
			EventID:     "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
			Speakers: []service.Speaker{{
				ID:       "9c08fbf8-160b-4a86-9981-aeddf4e3798e",
				Name:     "Manuel Iglesias",
				Bio:      "Manuel Iglesias draws from his life story when writing about the experiences of migrant workers. His first series of poems, My Father's Hands, appeared in The New Yorker and describes how his family crossed the Texas border to give Manuel and his brothers a better life. After receiving his Master's in Fine Arts from Columbia University, Manuel wrote three nonfiction novels about his experiences, including Under the Streaming Sun, which earned the National Prize for Arts and Sciences in 2008. Manuel is currently working on a collection of fictional short stories to be published in early 2021.",
				Headline: "CEO of Tesla",
				Photo:    "https://img.freepik.com/free-photo/handsome-young-businessman-shirt-eyeglasses_85574-6228.jpg?size=626&ext=jpg",
			}},
		}},
	}
	return event
}