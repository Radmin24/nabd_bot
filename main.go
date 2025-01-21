package main

import (
	"context"
	"log"
	"not_a_boring_date_bot/api"
	"not_a_boring_date_bot/bot"
	"not_a_boring_date_bot/cache"
	"not_a_boring_date_bot/config"
)

func main() {
	cfg := config.NewConfig()

	apiClient := api.NewClient(cfg.APIEndpoint, cfg.Timeout)

	redisCache := cache.NewCache(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	bot, err := bot.NewBot(cfg.TelegramToken, apiClient, redisCache, cfg.Debug)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	ctx := context.Background()

	go bot.CheckAPIStatus(bot, ctx, &cfg.APIEndpoint)

	log.Println("Bot started...")

	if err := bot.Start(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
