package usecases

import (
	"context"
	"log"

	"amadeus.tele.ai/repositories/api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramUC struct {
	Chai api.Restchai
	Bot  *tgbotapi.BotAPI
}

func NewTelegram(chai api.Restchai, botKey string, debug bool) *telegramUC {
	// initializing bot
	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	// debug bot
	bot.Debug = debug
	log.Printf("bot %s has been activated", bot.Self.UserName)

	return &telegramUC{
		Chai: chai,
		Bot:  bot,
	}
}

func (uc *telegramUC) Chat(ctx context.Context) {
	u := tgbotapi.NewUpdate(1)
	u.Timeout = 60

	updates := uc.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			reply := uc.Chai.GetChat(ctx, update.Message.Text)
			botMsg := tgbotapi.NewMessage(update.Message.From.ID, reply)
			uc.Bot.Send(botMsg)
		}
	}
}
