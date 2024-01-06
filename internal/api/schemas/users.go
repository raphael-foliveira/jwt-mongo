package schemas

import (
	"time"

	"github.com/raphael-foliveira/fiber-mongo/internal/api/models"
)

type CreateUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
}

type UpdateUser struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserDto struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func UserToDto(user *models.User) *UserDto {
	return &UserDto{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
