package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	RECRUITMENT_SERVICE_PORT string
	HTTP_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.RECRUITMENT_SERVICE_PORT = cast.ToString(coalesce("RECRUITMENT_SERVICE_PORT", ":5050"))
	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8081"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "staffer_db"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
