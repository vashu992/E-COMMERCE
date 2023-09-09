package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DBSet() *mongo.Client {

	connectionURI := "mongodb://localhost:27017"


	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()


	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB:", err)
		return nil
	}

	fmt.Println("Successfully connected to MongoDB")
	return client

}

var Client *mongo.Client = DBSet()


func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("E-Commerce").Collection(collectionName)
	return collection

}


func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	productCollection := client.Database("E-Commerce").Collection(collectionName)
	return productCollection
}
