package main

import (
	"context"
	"log"
	"net"
	"not_a_boring_date_bot/api"
	"not_a_boring_date_bot/bot"
	"not_a_boring_date_bot/cache"
	"not_a_boring_date_bot/config"
	pb "not_a_boring_date_bot/grpc"
	"not_a_boring_date_bot/service"

	"google.golang.org/grpc"
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

	grpcServer := grpc.NewServer()

	grpcService := &service.ServergRPC{
		Bot_g: bot,
	}

	pb.RegisterGRPCServiceServer(grpcServer, grpcService)

	listener, err := net.Listen("tcp", "127.0.0.1:"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Error listening on gRPC port: %v", err)
	}

	log.Println("gRPC server is running on port " + cfg.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	if err := bot.Start(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}

}
