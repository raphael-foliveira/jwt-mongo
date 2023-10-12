package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartMongo() (*mongo.Client, error) {
	mongoUrl := os.Getenv("MONGO_URL")
	return mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
}
