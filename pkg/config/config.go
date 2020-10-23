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

type TracingConfig struct {
	Enabled bool   `env:"TRACING_ENABLED" env-default:"false"`
	Host    string `env:"TRACING_HOST" env-default:"localhost"`
}

type ServiceConfig struct {
	API struct {
		Host string `env:"SERVICE_API_HOST" env-default:"localhost"`
		Port int `env:"SERVICE_API_PORT" env-default:"1111"`
	}
	Auth struct {
		Host string `env:"SERVICE_AUTH_HOST" env-default:"localhost"`
		Port int `env:"SERVICE_AUTH_PORT" env-default:"2222"`
	}
	Speaker struct {
		Host string `env:"SERVICE_SPEAKER_HOST" env-default:"localhost"`
		Port int `env:"SERVICE_SPEAKER_PORT" env-default:"3333"`
	}
	Event struct {
		Host string `env:"SERVICE_EVENT_HOST" env-default:"localhost"`
		Port int `env:"SERVICE_EVENT_PORT" env-default:"4444"`
	}
	Session struct {
		Host string `env:"SERVICE_SESSION_HOST" env-default:"localhost"`
		Port int `env:"SERVICE_SESSION_PORT" env-default:"5555"`
	}
}

type JwtConfig struct {
	SecretKey string `env:"JWT_SECRET_KEY" env-default:"V3RY_S3CR3T_K3Y"`
	CookieName string `env:"JWT_COOKIE_NAME" env-default:"token"`
}

type Config struct {
	Db      DbConfig
	Service ServiceConfig
	Tracing TracingConfig
	Jwt JwtConfig
}

// Read the environment variables into a `Config` struct.
func NewConfig() *Config {
	var c Config
	if err := cleanenv.ReadEnv(&c); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config has been loaded: %+v", c)
	return &c
}