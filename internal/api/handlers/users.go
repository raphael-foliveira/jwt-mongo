package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/service"
)

type Users interface {
	Create(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
}

type users struct {
	service service.Users
}

func NewUsersHandler(service service.Users) Users {
	return &users{service}
}

func (u *users) Create(c *fiber.Ctx) error {
	var createUserDto schemas.CreateUser
	if err := c.BodyParser(&createUserDto); err != nil {
		return err
	}
	result, err := u.service.Create(c.Context(), createUserDto)
	if err != nil {
		return &schemas.ApiErr{
			Code:    http.StatusConflict,
			Message: "User already exists",
		}
	}
	return c.Status(http.StatusCreated).JSON(result)
}

func (u *users) List(c *fiber.Ctx) error {
	users, err := u.service.List(c.Context())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(users)
}

func (u *users) Get(c *fiber.Ctx) error {
	user, err := u.service.Get(c.Context(), c.Params("id"))
	if err != nil {
		return &schemas.ApiErr{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (u *users) Delete(c *fiber.Ctx) error {
	if err := u.service.Delete(c.Context(), c.Params("id")); err != nil {
		return err
	}
	return c.SendStatus(http.StatusNoContent)
}

func (u *users) Update(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusMethodNotAllowed)
}

func (u *users) Login(c *fiber.Ctx) error {
	var loginDto schemas.UserLogin
	if err := c.BodyParser(&loginDto); err != nil {
		return err
	}
	token, err := u.service.Login(c.Context(), loginDto)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(token)
}

func (u *users) Authenticate(c *fiber.Ctx) error {
	var token schemas.ValidateToken
	if err := c.BodyParser(&token); err != nil {
		return err
	}
	tokenPayload, err := u.service.CheckToken(c.Context(), token.Token)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(tokenPayload)
}
