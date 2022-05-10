package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot     *tgbotapi.BotAPI
	ChatIDs []int64
}

func NewTelegramBot(botToken string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	return &TelegramBot{Bot: bot}, nil
}

func (b *TelegramBot) ListenForUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "suscribe":
			AddMemberChatID(update.Message.Chat.ID)
			fmt.Println(AllChatIDs)
			msg.Text = "Added"
		case "exchange":
			exchange, err := ScrapeExchange()

			if err != nil {
				fmt.Println(err)
				continue
			}

			msg.Text = exchange.GetFormattedMessage()
		default:
			msg.Text = "Unknown command"
		}

		if _, err := b.Bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func (b *TelegramBot) SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	b.Bot.Send(msg)
}
