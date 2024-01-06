package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/fiber-mongo/internal/api"
	"github.com/raphael-foliveira/fiber-mongo/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database := database.Start()
	defer func() {
		if err := database.Client().Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	if err := api.NewServer(database).Start(); err != nil {
		panic(err)
	}
}
