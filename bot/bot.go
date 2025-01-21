package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"not_a_boring_date_bot/api"
	"not_a_boring_date_bot/cache"
	"not_a_boring_date_bot/internal/models"
	"not_a_boring_date_bot/messages"
	"os/exec"
	"sort"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot   *tgbotapi.BotAPI
	api   *api.Client
	cache *cache.Cache
}

type CommandHandler func(ctx context.Context, update *tgbotapi.Update) error

func NewBot(token string, api *api.Client, cache *cache.Cache, debug bool) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = debug

	b := &Bot{
		bot:   bot,
		api:   api,
		cache: cache,
	}

	return b, nil
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	ctx := context.Background()
	for update := range updates {

		switch {
		case update.CallbackQuery != nil:
			if update.CallbackQuery != nil && update.CallbackQuery.Data == "yes_my_handler" {
				err := b.QueryMyHandler(ctx, update)
				if err != nil {
					log.Println(err)
				}
				return nil
			}
		default:
			err := b.sendFormController(ctx, update)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

func (b *Bot) sendFormController(ctx context.Context, update tgbotapi.Update) error {

	apiResp, err := b.api.SendCommand(ctx, update)
	if err != nil {
		if err := b.cache.SetAPIStatus(ctx, false); err != nil {
			log.Printf("Error setting API status: %v", err)
		}

		jsonData, err := json.Marshal(update)
		if err != nil {
			return err
		}

		if err := b.cache.AddUserToNotify(ctx, update.Message.Chat.ID, jsonData); err != nil {
			log.Printf("Error adding user to notify list: %v", err)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.APIUnavailable)
		_, err = b.bot.Send(msg)
		return err
	}

	// apiStatus, err := b.cache.GetAPIStatus(ctx)
	// if err != nil {
	// 	log.Printf("Error getting API status: %v", err)
	// }

	// if !apiStatus {
	// 	if err := b.notifyUsersAPIRestored(ctx); err != nil {
	// 		log.Printf("Error notifying users about API restoration: %v", err)
	// 	}
	// }

	var queue []*models.ControllerResponce

	queue = append(queue, apiResp)

	if apiResp.IsNextMsg {
		apiRespId, err := b.api.SendID(ctx, apiResp.Id)
		if err != nil {
			if err := b.cache.SetAPIStatus(ctx, false); err != nil {
				log.Printf("Error setting API status: %v", err)
			}

			jsonData, err := json.Marshal(update)
			if err != nil {
				return err
			}

			if err := b.cache.AddUserToNotify(ctx, update.Message.Chat.ID, jsonData); err != nil {
				log.Printf("Error adding user to notify list: %v", err)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.APIUnavailable)
			_, err = b.bot.Send(msg)
			return err
		}
		queue = append(queue, apiRespId)

		id := apiRespId.Id

		for {

			resp, err := b.api.SendID(ctx, id)
			if err != nil {
				if err := b.cache.SetAPIStatus(ctx, false); err != nil {
					log.Printf("Error setting API status: %v", err)
				}

				jsonData, err := json.Marshal(update)
				if err != nil {
					return err
				}

				if err := b.cache.AddUserToNotify(ctx, update.Message.Chat.ID, jsonData); err != nil {
					log.Printf("Error adding user to notify list: %v", err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.APIUnavailable)
				_, err = b.bot.Send(msg)
				return err
			}

			queue = append(queue, resp)
			if id == resp.Id {
				log.Println("Error lading next ID DUBLICATE:", id)
				break
			}

			id = resp.Id
			if !resp.IsNextMsg {
				break
			}
		}

	}

	if len(queue) > 0 {
		for _, v := range queue {
			Sender(v, b, update.Message.Chat.ID, ctx)
			if v.Delay > 0 {
				time.Sleep(time.Duration(v.Delay) * time.Second)
			}
		}
	} else {
		Sender(apiResp, b, update.Message.Chat.ID, ctx)
	}

	return nil
}

func (b *Bot) QueryMyHandler(ctx context.Context, update tgbotapi.Update) error {

	idstr := strconv.Itoa(int(update.CallbackQuery.From.ID))
	usersToNotify, err := b.cache.GetUsersToNotifyFromYES(ctx, idstr)

	fmt.Println("usersToNotify:", usersToNotify)

	if err != nil {
		log.Println(err)
	}

	var updatereturn *tgbotapi.Update
	err = json.Unmarshal([]byte(usersToNotify), &updatereturn)

	b.sendFormController(ctx, *updatereturn)

	return nil
}

func Sender(message *models.ControllerResponce, b *Bot, chatID int64, ctx context.Context) error {
	msg := tgbotapi.NewMessage(chatID, message.Answer)
	if message.IsKb {
		keyboard := message.Keyboard
		switch keyboard.Type {
		case "inline":
			msg.ReplyMarkup = generationInlineKeybord(keyboard.Button)
		case "reply":
			msg.ReplyMarkup = generationReplyKeybord(keyboard.Button)
		}
	} else {
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	_, err := b.bot.Send(msg)
	return err
}

func generationReplyKeybord(replayKeyboard []models.Button) tgbotapi.ReplyKeyboardMarkup {
	var keyboard tgbotapi.ReplyKeyboardMarkup

	sort.Slice(replayKeyboard, func(i, j int) bool {
		if replayKeyboard[i].Row == replayKeyboard[j].Row {
			return replayKeyboard[i].Order < replayKeyboard[j].Order
		}
		return replayKeyboard[i].Row < replayKeyboard[j].Row
	})

	var rowsMap = make(map[int][]tgbotapi.KeyboardButton)
	for _, button := range replayKeyboard {
		rowsMap[button.Row] = append(rowsMap[button.Row], tgbotapi.NewKeyboardButton(button.Caption))
	}

	for _, rowButtons := range rowsMap {
		keyboard.Keyboard = append(keyboard.Keyboard, rowButtons)
	}

	return keyboard
}

func generationInlineKeybord(inlineKeyboard []models.Button) tgbotapi.InlineKeyboardMarkup {
	sort.Slice(inlineKeyboard, func(i, j int) bool {
		if inlineKeyboard[i].Row == inlineKeyboard[j].Row {
			return inlineKeyboard[i].Order < inlineKeyboard[j].Order
		}
		return inlineKeyboard[i].Row < inlineKeyboard[j].Row
	})

	var rowsMap = make(map[int][]tgbotapi.InlineKeyboardButton)
	for _, button := range inlineKeyboard {
		btn := tgbotapi.NewInlineKeyboardButtonData(button.Caption, button.Data)
		rowsMap[button.Row] = append(rowsMap[button.Row], btn)
	}

	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, rowButtons := range rowsMap {
		keyboard = append(keyboard, rowButtons)
	}

	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

func (b *Bot) CheckAPIStatus(bot *Bot, ctx context.Context, urlApi *string) {
	for {
		apiStatus, err := bot.cache.GetAPIStatus(ctx)
		if err != nil {
			log.Println("Ошибка при получении статуса API:", err)
		}

		if apiStatus {
			usersToNotify, err := bot.cache.GetUsersToNotify(ctx)
			if err != nil {
				log.Println("Ошибка при получении пользователей для уведомления:", err)
				time.Sleep(10 * time.Second)
				continue
			}

			for id, mes := range usersToNotify {
				value, err := strconv.Atoi(id)
				if err != nil {
					log.Println("Ошибка при преобразовании userID:", err)
					continue
				}
				cmd := exec.Command("curl", *urlApi)
				_, err = cmd.CombinedOutput()
				if err != nil {
					if err := b.cache.SetAPIStatus(ctx, false); err != nil {
						log.Printf("Error setting API status: %v", err)
					}
				} else {
					if err := bot.SendNotificationRestore(int64(value)); err != nil {
						log.Println("Ошибка при отправке уведомления пользователю:", err)
					}
					err = bot.cache.ClearUsersToNotify(ctx, int64(value), mes)
					if err != nil {
						log.Println("Ошибка при отчистке списка уведомления пользователю:", err)
					}
				}
			}

		} else {
			cmd := exec.Command("curl", *urlApi)
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Println("Ошибка при выполнении запроса к API:", err)
			} else {
				if err := b.cache.SetAPIStatus(ctx, true); err != nil {
					log.Printf("Error setting API status: %v", err)
				}
			}
		}
		time.Sleep(10 * time.Second)
	}

}

func (b *Bot) SendNotificationRestore(ChatID int64) error {

	msg := tgbotapi.NewMessage(ChatID, messages.ServiceRestored)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", "yes_my_handler"),
			tgbotapi.NewInlineKeyboardButtonData("Нет", "no"),
		),
	)

	_, err := b.bot.Send(msg)
	return err

}
