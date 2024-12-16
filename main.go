package main

import (
	"log"
	"not_a_boring_date_bot/api"
	"not_a_boring_date_bot/bot"
	"not_a_boring_date_bot/cache"
	"not_a_boring_date_bot/config"
)

func main() {
	cfg := config.NewConfig()

	apiClient := api.NewClient(cfg.APIEndpoint)

	redisCache := cache.NewCache(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	bot, err := bot.NewBot(cfg.TelegramToken, apiClient, redisCache)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	log.Println("Bot started...")

	if err := bot.Start(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
