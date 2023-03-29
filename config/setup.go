package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalln(err)
	}

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("SepAppDB").Collection(collectionName)
}
