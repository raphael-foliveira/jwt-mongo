package service

import (
	"context"
	"net/http"

	"github.com/raphael-foliveira/fiber-mongo/internal/api/repository"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
)

var UsersService = &Users{}

type Users struct{}

func (u *Users) Create(c context.Context, dto schemas.CreateUser) (*schemas.UserDto, error) {
	user, err := repository.UsersRepository.Create(c, dto)
	if err != nil {
		return nil, err
	}
	return schemas.UserToDto(user), nil
}

func (u *Users) List(c context.Context) ([]schemas.UserDto, error) {
	users, err := repository.UsersRepository.List(c)
	if err != nil {
		return nil, err
	}
	usersView := []schemas.UserDto{}
	for _, user := range users {
		usersView = append(usersView, *schemas.UserToDto(&user))
	}
	return usersView, nil

}

func (u *Users) Get(c context.Context, id string) (*schemas.UserDto, error) {
	user, err := repository.UsersRepository.Get(c, id)
	if err != nil {
		return nil, err
	}
	return schemas.UserToDto(user), nil
}

func (u *Users) Delete(c context.Context, id string) error {
	return repository.UsersRepository.Delete(c, id)
}

func (u *Users) Login(c context.Context, dto schemas.UserLogin) (*schemas.LoginResponse, error) {
	user, err := repository.UsersRepository.GetByEmail(c, dto.Email)
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
	if _, err := repository.UsersRepository.Update(c, user.ID, &schemas.UpdateUser{
		Password: user.Password,
		Token:    user.Token,
	}); err != nil {
		return nil, err
	}
	return &schemas.LoginResponse{
		Token: token,
	}, nil
}

func (u *Users) CheckToken(c context.Context, token string) (*schemas.TokenPayload, error) {
	tokenPayload, err := jwtService.validateToken(token)
	if err != nil {
		return nil, err
	}
	user, err := repository.UsersRepository.GetByEmail(c, tokenPayload.Email)
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
