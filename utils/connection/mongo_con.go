package connection

import (
	"context"
	"log"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoCon(dbname string) *mongo.Database {
	ctx := context.Background()

	clientOpt := options.Client()
	clientOpt.ApplyURI("mongodb+srv://" + os.Getenv("MONGO_USERNAME") + ":" + url.QueryEscape(os.Getenv("MONGO_PASSWORD")) + "@cluster0.ovwnp4u.mongodb.net/?retryWrites=true&w=majority")

	cli, err := mongo.NewClient(clientOpt)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = cli.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cli.Database(dbname)
}
