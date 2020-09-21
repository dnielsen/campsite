package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"path/filepath"
	"runtime"
	"strings"
)

type DbConfig struct {
	Name string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
	User string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:"postgres"`
	// When you run the program locally, you need the host to be `localhost`
	// When it's run with docker-compose, it needs to be `cassandra`, so we're
	// passing it in the `docker.env` file.
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	SSLMode     string `yaml:"sslmode" env:"DB_SSLMODE" env-default:"disable"`
}

type Config struct {
	Db DbConfig `yaml:"db"`
	Server struct {
		Address string `yaml:"address" env:"SERVER_ADDRESS" env-default:"0.0.0.0:5555"`
	} `yaml:"server"`
	Service struct {
		Speaker struct{
			Address string `yaml:"address" env:"SERVICE_SPEAKER_ADDRESS" env-default:"localhost:5555"`
		} `yaml:"speaker"`
	} `yaml:"service"`
}

type SmtpConfig struct {
	Address  string `yaml:"address" env:"SMTP_ADDRESS"`
	From     string `yaml:"from" env:"SMTP_FROM"`
	Password string `yaml:"password" env:"SMTP_PASSWORD"`
}

// Try to read variables from the config file.
// If it fails, read them from environment.
func GetConfig(filename string) (*Config, error) {
	var c Config
	path := getConfigPath(filename)
	if err := cleanenv.ReadConfig(path, &c); err != nil {
		if err := cleanenv.ReadEnv(&c); err != nil {
			return nil, err
		}
	}
	return &c, nil
}

// Return the path on disk to the configs
func getConfigPath(configFilename string) string {
	_, currentFilename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Join(filepath.Dir(currentFilename), "../../configs/", configFilename)
}

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