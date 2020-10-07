package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type DbConfig struct {
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	SSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

type ServerConfig struct {
	Tracing TracingConfig
	Port string `env:"SERVER_PORT" env-default:"4444"`
}

type TracingConfig struct {
	Enabled bool `env:"SERVER_TRACING_ENABLED" env-default:"false"`
	Host string `env:"SERVER_TRACING_HOST" env-default:"localhost"`
}

type Config struct {
	Db     DbConfig
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
