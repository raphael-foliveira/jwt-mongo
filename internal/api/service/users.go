package service

import (
	"context"
	"net/http"

	"github.com/raphael-foliveira/fiber-mongo/internal/api/repository"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
)

type Users struct {
	repo repository.Users
}

func NewUsers(repo repository.Users) *Users {
	return &Users{repo}
}

func (u *Users) Create(c context.Context, dto schemas.CreateUser) (*schemas.UserView, error) {
	return u.repo.Create(c, dto)
}

func (u *Users) List(c context.Context) ([]schemas.UserView, error) {
	return u.repo.List(c)
}

func (u *Users) Get(c context.Context, id string) (*schemas.UserView, error) {
	return u.repo.Get(c, id)
}

func (u *Users) Delete(c context.Context, id string) error {
	return u.repo.Delete(c, id)
}

func (u *Users) Login(c context.Context, dto schemas.UserLogin) (*schemas.LoginResponse, error) {
	user, err := u.repo.GetByEmail(c, dto.Email)
	if err != nil {
		return nil, &schemas.ApiErr{
			Code:    http.StatusUnauthorized,
			Message: "invalid credentials",
		}
	}
	if user.Password != dto.Password {
		return nil, &schemas.ApiErr{
			Code:    http.StatusUnauthorized,
			Message: "invalid credentials",
		}
	}
	token, err := generateToken(schemas.TokenPayload{
		Sub:   user.ID,
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}
	return &schemas.LoginResponse{
		Token: token,
	}, nil
}

func (u *Users) CheckToken(c context.Context, token string) (*schemas.TokenPayload, error) {
	return validateToken(token)
}
