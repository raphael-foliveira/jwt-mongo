package handlers

import "github.com/gofiber/fiber/v2"

type HealthCheck struct{}

func (hc *HealthCheck) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func NewHealthCheckHandler() *HealthCheck {
	return &HealthCheck{}
}
