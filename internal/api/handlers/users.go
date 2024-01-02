package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/service"
)

var UsersHandler = &Users{}

type Users struct{}

func (u *Users) Create(c *fiber.Ctx) error {
	var createUserDto schemas.CreateUser
	if err := c.BodyParser(&createUserDto); err != nil {
		return err
	}
	result, err := service.UsersService.Create(c.Context(), createUserDto)
	if err != nil {
		return &schemas.ApiErr{
			Code:    http.StatusConflict,
			Message: "User already exists",
		}
	}
	return c.Status(http.StatusCreated).JSON(result)
}

func (u *Users) List(c *fiber.Ctx) error {
	users, err := service.UsersService.List(c.Context())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(users)
}

func (u *Users) Get(c *fiber.Ctx) error {
	user, err := service.UsersService.Get(c.Context(), c.Params("id"))
	if err != nil {
		return &schemas.ApiErr{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (u *Users) Delete(c *fiber.Ctx) error {
	if err := service.UsersService.Delete(c.Context(), c.Params("id")); err != nil {
		return err
	}
	return c.SendStatus(http.StatusNoContent)
}

func (u *Users) Update(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusMethodNotAllowed)
}

func (u *Users) Login(c *fiber.Ctx) error {
	var loginDto schemas.UserLogin
	if err := c.BodyParser(&loginDto); err != nil {
		return err
	}
	token, err := service.UsersService.Login(c.Context(), loginDto)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(token)
}

func (u *Users) Authenticate(c *fiber.Ctx) error {
	var token schemas.ValidateToken
	if err := c.BodyParser(&token); err != nil {
		return err
	}
	tokenPayload, err := service.UsersService.CheckToken(c.Context(), token.Token)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(tokenPayload)
}
