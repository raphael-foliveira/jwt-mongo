package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repositories struct {
	Users Users
}

func StartRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users: NewUsersRepository(db.Collection("users")),
	}
}
