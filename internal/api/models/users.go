package models

import "time"

type User struct {
	ID        string    `bson:"_id" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}
