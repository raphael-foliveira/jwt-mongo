package service

import (
	"context"
	"net/http"

	"github.com/raphael-foliveira/fiber-mongo/internal/api/repository"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
)

type Users interface {
	Create(c context.Context, dto schemas.CreateUser) (*schemas.UserDto, error)
	List(c context.Context) ([]schemas.UserDto, error)
	Get(c context.Context, id string) (*schemas.UserDto, error)
	Delete(c context.Context, id string) error
	Login(c context.Context, dto schemas.UserLogin) (*schemas.LoginResponse, error)
	CheckToken(c context.Context, token string) (*schemas.TokenPayload, error)
}

type users struct {
	repo repository.Users
}

func NewUsersService(repository repository.Users) Users {
	return &users{repository}
}

func (u *users) Create(c context.Context, dto schemas.CreateUser) (*schemas.UserDto, error) {
	user, err := u.repo.Create(c, dto)
	if err != nil {
		return nil, err
	}
	return schemas.UserToDto(user), nil
}

func (u *users) List(c context.Context) ([]schemas.UserDto, error) {
	users, err := u.repo.List(c)
	if err != nil {
		return nil, err
	}
	usersView := []schemas.UserDto{}
	for _, user := range users {
		usersView = append(usersView, *schemas.UserToDto(&user))
	}
	return usersView, nil

}

func (u *users) Get(c context.Context, id string) (*schemas.UserDto, error) {
	user, err := u.repo.Get(c, id)
	if err != nil {
		return nil, err
	}
	return schemas.UserToDto(user), nil
}

func (u *users) Delete(c context.Context, id string) error {
	return u.repo.Delete(c, id)
}

func (u *users) Login(c context.Context, dto schemas.UserLogin) (*schemas.LoginResponse, error) {
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
	token, err := jwtService.generateToken(schemas.TokenPayload{
		Sub:   user.ID,
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}
	user.Token = token
	if _, err := u.repo.Update(c, user.ID, &schemas.UpdateUser{
		Password: user.Password,
		Token:    user.Token,
	}); err != nil {
		return nil, err
	}
	return &schemas.LoginResponse{
		Token: token,
	}, nil
}

func (u *users) CheckToken(c context.Context, token string) (*schemas.TokenPayload, error) {
	tokenPayload, err := jwtService.validateToken(token)
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
