package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv string
	DB     DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

func Load() (Config, error) {
	appEnv, err := getEnv("APP_ENV")
	if err != nil {
		return Config{}, err
	}

	host, err := getEnv("DB_HOST")
	if err != nil {
		return Config{}, err
	}

	port, err := getEnv("DB_PORT")
	if err != nil {
		return Config{}, err
	}

	user, err := getEnv("DB_USER")
	if err != nil {
		return Config{}, err
	}

	password, err := getEnv("DB_PASSWORD")
	if err != nil {
		return Config{}, err
	}

	name, err := getEnv("DB_NAME")
	if err != nil {
		return Config{}, err
	}

	sslMode, err := getEnv("DB_SSLMODE")
	if err != nil {
		return Config{}, err
	}

	timeZone, err := getEnv("DB_TIMEZONE")
	if err != nil {
		return Config{}, err
	}

	return Config{
		AppEnv: appEnv,
		DB: DBConfig{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Name:     name,
			SSLMode:  sslMode,
			TimeZone: timeZone,
		},
	}, nil
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name,
		c.SSLMode,
		c.TimeZone,
	)
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}

	return value, nil
}
