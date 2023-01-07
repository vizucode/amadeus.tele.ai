package main

import (
	"context"
	"log"
	"os"
	"sync"

	"amadeus.tele.ai/repositories/api"
	"amadeus.tele.ai/repositories/localstorage"
	uc "amadeus.tele.ai/usecases"
	"github.com/joho/godotenv"
)

func main() {
	wg := new(sync.WaitGroup)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	// initializing chai
	chai := api.NewchaiML(os.Getenv("DEV_KEY"), os.Getenv("DEV_UID"), os.Getenv("URL_TARGET"))

	// initializing localstorage

	localstorage := localstorage.NewInternalFile()

	// initializing usecase bot
	teleBot := uc.NewTelegram(chai, localstorage, os.Getenv("BOT_TELE_API_KEY"), false)

	// starting the bot
	wg.Add(1)
	go start(teleBot, wg)
	wg.Wait()
}

func start(chat uc.PlatformUC, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx := context.TODO()
	chat.Chat(ctx)
}
