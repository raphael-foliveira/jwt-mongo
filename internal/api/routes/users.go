package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/handlers"
)

func Users(app *fiber.App, handler handlers.Users) {
	log.Println("setting up users routes")
	app.Route("/users", func(router fiber.Router) {
		router.Get("", handler.List)
		router.Post("", handler.Create)
		router.Post("/login", handler.Login)
		router.Post("/check-token", handler.Authenticate)
		router.Get("/:id", handler.Get)
		router.Delete("/:id", handler.Delete)
	})
}
