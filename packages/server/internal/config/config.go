package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"path/filepath"
	"runtime"
)

type DbConfig struct {
	Name string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
	User string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:"postgres"`
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	SSLMode     string `yaml:"sslmode" env:"DB_SSLMODE" env-default:"disable"`
}

type ServerConfig struct {
	Port string `yaml:"port" env:"SERVER_PORT" env-default:"4444"`
}

type Config struct {
	Db DbConfig `yaml:"db"`
	Server ServerConfig `yaml:"server"`
}

// If the filename isn't an empty string read the config from configs directory
// which is located in the project's root directory.
// Else, read the variables from the environment.
func getConfig(filename string) (*Config, error) {
	var c Config
	if filename != "" {
		// Read the config from the configs/{filename} file.
		// For example: configs/development.yml
		path := getConfigPath(filename)
		if err := cleanenv.ReadConfig(path, &c); err != nil {
			return nil, err
		}
	} else {
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

// Initialize the config. It will throw if an error occurs.
func NewConfig(filename string) *Config {
	c, err := getConfig(filename)
	if err != nil {
		log.Fatalf("Failed to load config: %v",err)
	}

	log.Printf("Config has been loaded: %v", c)
	return c
}