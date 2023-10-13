package handlers

import "github.com/gofiber/fiber/v2"

func parseIdFromParams(c *fiber.Ctx) string {
	return c.Params("id")
}


