package usecases

import (
	"context"
	"log"

	"amadeus.tele.ai/repositories/api"
	"amadeus.tele.ai/repositories/localstorage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramUC struct {
	Chai api.Restchai
	Ls   localstorage.Localstorage
	Bot  *tgbotapi.BotAPI
}

func NewTelegram(chai api.Restchai, ls localstorage.Localstorage, botKey string, debug bool) *telegramUC {
	// initializing bot
	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	// debug bot
	bot.Debug = debug
	log.Printf("%s has been activated", bot.Self.UserName)

	return &telegramUC{
		Chai: chai,
		Ls:   ls,
		Bot:  bot,
	}
}

func (uc *telegramUC) Chat(ctx context.Context) {
	u := tgbotapi.NewUpdate(1)
	u.Timeout = 60

	updates := uc.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			// read
			msg := uc.Ls.Read("memory.json")
			tempMsg := msg + "\nMe:" + update.Message.Text + "\nKuristina:"
			reply := uc.Chai.GetChat(ctx, tempMsg)
			if reply == "" {
				reply = "i'm stuck.."
			}

			// write
			msg = msg + "\nMe:" + update.Message.Text + "\nKuristina:" + reply
			uc.Ls.Write("memory.json", msg)

			botMsg := tgbotapi.NewMessage(update.Message.From.ID, reply)
			uc.Bot.Send(botMsg)
		}
	}
}
