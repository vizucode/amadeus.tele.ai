package usecases

import (
	"context"
	"log"

	"amadeus.tele.ai/repositories/api"
	"amadeus.tele.ai/repositories/localstorage"

	"amadeus.tele.ai/utils/translate"
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

			// translate
			translatedText, err := translate.Translate(update.Message.Text, "auto", "en")

			if err != nil {
				log.Fatal(err.Error())
			}

			tempMsg := msg + "\nMe:" + translatedText + "\nKuristina:"
			reply := uc.Chai.GetChat(ctx, tempMsg)
			if reply == "" {
				reply = "i'm stuck.."
			}

			// write
			msg = msg + "\nMe:" + translatedText + "\nKuristina:" + reply
			uc.Ls.Write("memory.json", msg)

			// tranlated reply
			translatedReply, err := translate.Translate(reply, "en", "id")
			if err != nil {
				log.Fatal(err.Error())
			}

			botMsg := tgbotapi.NewMessage(update.Message.From.ID, translatedReply)
			uc.Bot.Send(botMsg)
		}
	}
}
