package handlers

import "github.com/gofiber/fiber/v2"

func parseObjectIdFromParams(c *fiber.Ctx) string {
	return c.Params("id")
}


