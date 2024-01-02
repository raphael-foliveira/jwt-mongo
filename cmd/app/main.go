package main

import (
	"context"

	"github.com/raphael-foliveira/fiber-mongo/internal/api"
	"github.com/raphael-foliveira/fiber-mongo/internal/database"
)

func main() {
	defer func() {
		if err := database.MongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	if err := api.NewServer().Start(); err != nil {
		panic(err)
	}
}
