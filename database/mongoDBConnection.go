package database

import (
	"context"
	"flag"
	"time"

	"github.com/Som-Kesharwani/shared-service/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	Client = DBInstance()
}

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
		return nil
	}
	logger.Info.Println("Connected Successfully to Database!!!")
	return client

}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	logger.Info.Printf("Try Mongo DB Connection")
	var collection *mongo.Collection = Client.Database("Test").Collection(collectionName)
	logger.Info.Printf("Connection Successfull!!")
	return collection
}
