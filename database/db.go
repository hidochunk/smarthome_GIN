package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"smarthome_GIN/config"
)

var DB *mongo.Client = Init()

func Init() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.EnvMongoURI()))
	if err != nil {
		log.Fatal("Error connecting mongodb")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error testing connection")
	}

	log.Println("Finish database Init process")

	return client
}
