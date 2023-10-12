package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/handlers"
)

func Users(app *fiber.App, h *handlers.Users) {
	log.Println("setting up users routes")
	app.Route("/users", func(router fiber.Router) {
		router.Get("", h.List)
		router.Post("", h.Create)
		router.Post("/login", h.Login)
		router.Post("/check-token", h.Authenticate)
		router.Get("/:id", h.Get)
		router.Delete("/:id", h.Delete)
	})
}
