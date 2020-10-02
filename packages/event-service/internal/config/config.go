
package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
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
	Tracing TracingConfig
	Port string `env:"SERVICE_EVENT_PORT" env-default:"4444"`
}

type TracingConfig struct {
	Enabled bool `env:"SERVER_TRACING_ENABLED" env-default:"false"`
	Host string `env:"SERVER_TRACING_HOST" env-default:"localhost"`
}

type ServiceConfig struct {
	Speaker struct{
		Host string `env:"SERVICE_SPEAKER_HOST" env-default:"localhost"`
		Port string `env:"SERVICE_SPEAKER_PORT" env-default:"3333"`
	}
	Session struct{
		Host string `env:"SERVICE_SESSION_HOST" env-default:"localhost"`
		Port string `env:"SERVICE_SESSION_PORT" env-default:"5555"`
	}
}

type Config struct {
	Db DbConfig
	Server ServerConfig
	Service ServiceConfig
}

// Read the environment variables into a `Config` struct.
func NewConfig() *Config {
	var c Config
	if err := cleanenv.ReadEnv(&c); err != nil {
		log.Fatalf("Failed to load config: %v",err)
	}

	log.Printf("Config has been loaded: %+v", c)
	return &c
}

