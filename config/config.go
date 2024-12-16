package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	APIEndpoint   string
	RedisAddr     string
	RedisPassword string
	VersionAPI    string
	RedisDB       int
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		log.Printf("Warning: Invalid REDIS_DB value, using default: 0")
		redisDB = 0
	}

	conf := &Config{
		TelegramToken: getEnv("TELEGRAM_TOKEN", ""),
		APIEndpoint:   getEnv("API_ENDPOINT", "http://127.0.0.1:8080/botapi/"),
		RedisAddr:     getEnv("REDIS_ADDR", "127.0.0.1:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		VersionAPI:    getEnv("VERSION_API", "v1"),
		RedisDB:       redisDB,
	}

	conf.APIEndpoint = conf.APIEndpoint + conf.VersionAPI + "/"

	log.Println("Base API endpoint:", conf.APIEndpoint)

	return conf
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
