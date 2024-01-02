package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/fiber-mongo/internal/api"
	"github.com/raphael-foliveira/fiber-mongo/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	if err := database.StartMongo(); err != nil {
		panic(err)
	}
	defer func() {
		if err := database.MongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	if err := api.NewServer().Start(); err != nil {
		panic(err)
	}
}
