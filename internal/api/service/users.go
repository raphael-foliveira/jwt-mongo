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
	user, err := u.repo.Create(c, dto)
	if err != nil {
		return nil, err
	}
	return &schemas.UserView{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *Users) List(c context.Context) ([]schemas.UserView, error) {
	users, err := u.repo.List(c)
	if err != nil {
		return nil, err
	}
	usersView := []schemas.UserView{}
	for _, user := range users {
		usersView = append(usersView, schemas.UserView{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}
	return usersView, nil

}

func (u *Users) Get(c context.Context, id string) (*schemas.UserView, error) {
	user, err := u.repo.Get(c, id)
	if err != nil {
		return nil, err
	}
	return &schemas.UserView{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
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
	user.Token = token
	if _, err := u.repo.Update(c, user); err != nil {
		return nil, err
	}
	return &schemas.LoginResponse{
		Token: token,
	}, nil
}

func (u *Users) CheckToken(c context.Context, token string) (*schemas.TokenPayload, error) {
	tokenPayload, err := validateToken(token)
	if err != nil {
		return nil, err
	}
	user, err := u.repo.GetByEmail(c, tokenPayload.Email)
	if err != nil {
		return nil, err
	}
	if user.Token != token {
		return nil, &schemas.ApiErr{
			Code:    http.StatusUnauthorized,
			Message: "invalid token",
		}
	}
	return tokenPayload, nil
}
