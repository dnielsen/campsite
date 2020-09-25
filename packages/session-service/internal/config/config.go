package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"strings"
)

type DbConfig struct {
	Name string `env:"DB_NAME" env-default:"postgres"`
	User string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	SSLMode     string `env:"DB_SSLMODE" env-default:"disable"`
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT" env-default:"5555"`
}

type Config struct {
	Db DbConfig
	Server ServerConfig
}

// Read the environment variables into a `Config` struct.
func NewConfig() *Config {
	var c Config
	if err := cleanenv.ReadEnv(&c); err != nil {
		log.Fatalf("Failed to load config: %v",err)
	}

	log.Printf("Config has been loaded: %v", c)
	return &c
}


// Example result: "user=postgres password=postgres sslmode=disable ..."
func GetDbConnString(c *DbConfig) string  {
	vals := getDbValues(c)
	var p []string
	for k, v := range vals {
		p = append(p, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(p, " ")
}

func getDbValues(c *DbConfig) map[string]string {
	p := map[string]string{}
	setIfNotEmpty(p, "dbname", c.Name)
	setIfNotEmpty(p, "host", c.Host)
	setIfNotEmpty(p, "user", c.User)
	setIfNotEmpty(p, "password", c.Password)
	setIfNotEmpty(p, "port", c.Port)
	setIfNotEmpty(p, "sslmode", c.SSLMode)
	return p
}

func setIfNotEmpty(m map[string]string, key, val string) {
	if val != "" {
		m[key] = val
	}
}
