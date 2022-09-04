package modules

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection *mongo.Collection
	ctx            = context.TODO()
)

func InitDB() {
	host := os.Getenv("DatabaseHost")
	port := os.Getenv("DatabasePort")

	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port + "/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	userCollection = client.Database("go-tutor").Collection("users")
}

func UserCollection() (*mongo.Collection, context.Context) {
	return userCollection, ctx
}
