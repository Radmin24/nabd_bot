package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBot   string
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

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil || redisDB < 0 {
		redisDB = 0
	} else if err != nil {
		log.Printf("Warning: Invalid REDIS_DB value, using default: %d", 0)
		redisDB = 0
	}

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil || timeout < 0 {
		timeout = 10
	} else if err != nil {
		log.Printf("Warning: Invalid TIMEOUT value, using default: %d", 10)
		timeout = 10
	}

	gRPCtimeout, err := strconv.Atoi(os.Getenv("GRPCTIMEOUT"))
	if err != nil || gRPCtimeout < 0 {
		gRPCtimeout = 10
	} else if err != nil {
		log.Printf("Warning: Invalid GRPCTIMEOUT value, using default: %d", 10)
		gRPCtimeout = 10
	}

	conf := &Config{
		TelegramBot:   os.Getenv("TELEGRAM_BOT"),
		APIEndpoint:   os.Getenv("API_ENDPOINT"),
		APIorm:        os.Getenv("API_ORM_URL"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		VersionAPI:    os.Getenv("VERSION_API"),
		Timeout:       timeout,
		RedisDB:       redisDB,
		GRPCPort:      os.Getenv("GRPC_PORT"),
		GRPCTimeout:   gRPCtimeout,
	}

	if conf.TelegramBot == "" {
		log.Fatal("Error: TELEGRAM_BOT is not set in env  compose")
	}

	if conf.APIEndpoint == "" {
		log.Fatal("Error: API_ENDPOINT is not set in env  compose")
	}
	if conf.RedisAddr == "" {
		log.Fatal("Error: REDIS_ADDR is not set in env  compose")
	}

	if conf.RedisPassword == "" {
		log.Println("Use default Redis password ")
	}

	if conf.APIorm == "" {
		log.Println("Use default API orm url http://127.0.0.1:8081/")
	}

	if conf.VersionAPI == "" {
		log.Println("Use default version API v1")
	}

	Debug := os.Getenv("DEBUG")

	if Debug == "" {
		log.Println("Use default debug mode folse")
	} else if Debug == "true" {
		log.Println("Use debug mode true")
		conf.Debug = true
	} else {
		conf.Debug = false
	}

	if conf.GRPCPort == "" {
		log.Println("Use default GRPC port 50051")
		conf.GRPCPort = "50051"
	}

	conf.APIorm = conf.APIorm + "botconfig/"

	conf.TelegramToken = GetTocken(conf.TelegramBot, conf.APIorm)

	conf.APIEndpoint = conf.APIEndpoint + conf.VersionAPI + "/"

	log.Println("Config start:", conf)

	return conf
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetTocken(botName string, urlorm string) string {

	url := urlorm + "?caption=" + botName

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	stingtocken := string(body)
	tocken := strings.Trim(stingtocken, "\"")

	fmt.Println("Use tocken:", tocken)

	return tocken
}
