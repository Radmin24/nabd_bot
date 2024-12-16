package bot

import (
	"context"
	"log"
	"not_a_boring_date_bot/api"
	"not_a_boring_date_bot/cache"
	"not_a_boring_date_bot/messages"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	api      *api.Client
	cache    *cache.Cache
	commands map[string]CommandHandler
}

type CommandHandler func(ctx context.Context, update *tgbotapi.Update) error

func NewBot(token string, api *api.Client, cache *cache.Cache) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	b := &Bot{
		bot:   bot,
		api:   api,
		cache: cache,
	}

	b.commands = map[string]CommandHandler{
		"start": b.handleCommand,
	}

	return b, nil
}

func (b *Bot) Start() error {
	// Init
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	// Создайте новый контекст для бота, это необходимо для правильной работы кэша и API-клиента
	ctx := context.Background()
	for update := range updates {
		if update.Message == nil {
			continue // пока скипаем
		}

		if update.Message.IsCommand() {
			command := update.Message.Command()
			if handler, ok := b.commands[command]; ok {
				go func(ctx context.Context, update tgbotapi.Update) {
					if err := handler(ctx, &update); err != nil {
						log.Printf("Error handling command %s: %v", command, err)
					}
				}(ctx, update)
			}
		}
	}

	return nil
}

func (b *Bot) handleCommand(ctx context.Context, update *tgbotapi.Update) error {

	apiResp, err := b.api.SendCommand(ctx, update)
	if err != nil {
		if err := b.cache.SetAPIStatus(ctx, false); err != nil {
			log.Printf("Error setting API status: %v", err)
		}
		if err := b.cache.AddUserToNotify(ctx, update.Message.Chat.ID); err != nil {
			log.Printf("Error adding user to notify list: %v", err)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.APIUnavailable)
		_, err := b.bot.Send(msg)
		return err
	}

	apiStatus, err := b.cache.GetAPIStatus(ctx)
	if err != nil {
		log.Printf("Error getting API status: %v", err)
	}

	if !apiStatus {
		if err := b.notifyUsersAPIRestored(ctx); err != nil {
			log.Printf("Error notifying users about API restoration: %v", err)
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, apiResp.Message)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) notifyUsersAPIRestored(ctx context.Context) error {
	users, err := b.cache.GetUsersToNotify(ctx)
	if err != nil {
		return err
	}

	for _, userStr := range users {
		userID, err := strconv.ParseInt(userStr, 10, 64)
		if err != nil {
			log.Printf("Error parsing user ID: %v", err)
			continue
		}

		msg := tgbotapi.NewMessage(userID, messages.ServiceRestored)
		if _, err := b.bot.Send(msg); err != nil {
			log.Printf("Error sending restoration notification to user %d: %v", userID, err)
		}
	}

	if err := b.cache.ClearUsersToNotify(ctx); err != nil {
		return err
	}

	return b.cache.SetAPIStatus(ctx, true)
}
