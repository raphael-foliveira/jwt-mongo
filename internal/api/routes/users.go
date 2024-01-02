package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/handlers"
)

func Users(app *fiber.App) {
	log.Println("setting up users routes")
	app.Route("/users", func(router fiber.Router) {
		router.Get("", handlers.UsersHandler.List)
		router.Post("", handlers.UsersHandler.Create)
		router.Post("/login", handlers.UsersHandler.Login)
		router.Post("/check-token", handlers.UsersHandler.Authenticate)
		router.Get("/:id", handlers.UsersHandler.Get)
		router.Delete("/:id", handlers.UsersHandler.Delete)
	})
}
