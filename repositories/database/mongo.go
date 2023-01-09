package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDB struct {
	Mongo *mongo.Database
}

func NewMongo(client *mongo.Database) *mongoDB {
	return &mongoDB{
		Mongo: client,
	}
}

func (db *mongoDB) Read(fileName string) string {
	ctx := context.TODO()
	where := bson.M{}

	csr := db.Mongo.Collection(fileName).FindOne(ctx, where)

	body := struct {
		Data string
	}{}

	err := csr.Decode(&body)
	if err != nil {
		log.Fatal(err.Error())
	}

	return body.Data
}

func (db *mongoDB) Write(fileName string, data string) {
	ctx := context.TODO()

	// bson struct
	body := struct {
		Data string `bson:"data"`
	}{
		Data: data,
	}

	_, err := db.Mongo.Collection(fileName).UpdateOne(ctx, bson.M{"unique_id": 0}, bson.M{"$set": body})
	if err != nil {
		log.Fatal(err.Error())
	}
}
