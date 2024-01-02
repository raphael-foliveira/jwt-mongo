package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient     *mongo.Client
	MongoDb         *mongo.Database
	UsersCollection *mongo.Collection
)

func StartMongo() (err error) {
	mongoUrl := os.Getenv("MONGO_URL")
	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return err
	}
	MongoDb = MongoClient.Database("fibermongo")
	getCollections()
	return createIndexes()
}

func createIndexes() (err error) {
	_, err = UsersCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "email", Value: 1},
				{Key: "username", Value: 1},
			},
			Options: options.Index().SetUnique(true)})
	return err
}

func getCollections() (err error) {
	UsersCollection = MongoDb.Collection("users")
	return nil
}
