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
	APIorm        string
	RedisAddr     string
	RedisPassword string
	VersionAPI    string
	Timeout       int
	Debug         bool
	RedisDB       int
	GRPCPort      string
	GRPCTimeout   int
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

	timeout, err := strconv.Atoi(getEnv("TIMEOUT", "10"))
	if err != nil {
		log.Printf("Warning: Invalid TIMEOUT value, using default: 10")
		timeout = 10
	}

	gRPCtimeout, err := strconv.Atoi(getEnv("GRPCTIMEOUT", "10"))
	if err != nil {
		log.Printf("Warning: Invalid GRPCTIMEOUT value, using default: 10")
		timeout = 10
	}

	conf := &Config{
		TelegramToken: getEnv("TELEGRAM_TOKEN", ""),
		APIEndpoint:   getEnv("API_ENDPOINT", "http://127.0.0.1:8080/botapi/"),
		APIorm:        getEnv("API_ORM_URL", "http://127.0.0.1:8081/"),
		RedisAddr:     getEnv("REDIS_ADDR", "127.0.0.1:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		VersionAPI:    getEnv("VERSION_API", "v1"),
		Timeout:       timeout,
		Debug:         getEnv("DEBUG", "false") == "true",
		RedisDB:       redisDB,
		GRPCPort:      getEnv("GRPC_PORT", "50051"),
		GRPCTimeout:   gRPCtimeout,
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
