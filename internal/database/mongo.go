package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartMongo() (*mongo.Client, error) {
	mongoUrl := os.Getenv("MONGO_URL")
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	return createIndexes(mongoClient)
}

func createIndexes(mongoClient *mongo.Client) (*mongo.Client, error) {
	db := mongoClient.Database("fibermongo")
	_, err := db.Collection("users").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "email", Value: 1},
				{Key: "username", Value: 1},
			},
			Options: options.Index().SetUnique(true)})
	return mongoClient, err
}
