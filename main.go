package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

var (
	BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")
)

const MEMBERS_FILE_PATH = "./members.txt"

func main() {
	AllChatIDs, err := LoadMembersChatIDs()
	if err != nil {
		fmt.Println(err)
		panic("Cannot load members chat ids")
	}

	fmt.Println(BOT_TOKEN)
	bot, err := NewTelegramBot(BOT_TOKEN)

	if err != nil {
		fmt.Println("error creating telegram bot")
		panic(err)
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(4).Hours().Do(func() {
		exchange, nil := ScrapeExchange()

		if err == nil {
			for _, chatID := range AllChatIDs {
				bot.SendMessage(chatID, exchange.GetFormattedMessage())
			}
		}

	})

	s.StartAsync()

	bot.ListenForUpdates()
}
