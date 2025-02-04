package database

import (
	"context"
	"flag"
	"time"

	"github.com/Som-Kesharwani/shared-service/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	MongoDB := flag.String("mongo", "mongodb://localhost:27017", "MongoDB connection string")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*MongoDB))

	if err != nil {
		logger.Error.Printf("Failed to connect DataBase. Error :%s", err)
		return nil
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logger.Error.Printf("Failed to ping DataBase. Error :%s", err)
	}
	logger.Info.Println("Connected Successfully to Database!!!")
	return client

}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("Test").Collection(collectionName)
	return collection
}
