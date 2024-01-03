package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
)

func Start() {
	mongoUrl := os.Getenv("MONGO_URL")
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}
	createIndexes(mongoClient.Database("fibermongo"))
	MongoClient = mongoClient
}

func createIndexes(database *mongo.Database) (err error) {
	_, err = database.Collection("users").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "email", Value: 1},
				{Key: "username", Value: 1},
			},
			Options: options.Index().SetUnique(true)})
	return err
}
