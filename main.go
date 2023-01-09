package main

import (
	"context"
	"log"
	"os"

	"amadeus.tele.ai/repositories/api"
	"amadeus.tele.ai/repositories/database"
	uc "amadeus.tele.ai/usecases"
	mongoCLI "amadeus.tele.ai/utils/connection"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// initializing chai
	chai := api.NewchaiML(os.Getenv("DEV_KEY"), os.Getenv("DEV_UID"), os.Getenv("URL_TARGET"))

	// initializing database
	mongoCli := mongoCLI.MongoCon(os.Getenv("MONGO_DBNAME"))
	db := database.NewMongo(mongoCli)

	// initializing usecase bot
	teleBot := uc.NewTelegram(chai, db, os.Getenv("BOT_TELE_API_KEY"), false)

	// starting the bot
	start(teleBot)
}

func start(chat uc.PlatformUC) {
	ctx := context.TODO()
	chat.Chat(ctx)
}
