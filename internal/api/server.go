package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/handlers"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/repository"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/routes"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/service"
	"github.com/raphael-foliveira/fiber-mongo/internal/database"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: apiErrorHandler,
	})
	return &Server{app}
}

func (s *Server) Start() error {
	s.app.Use(cors.New())
	s.app.Use(logger.New())

	repositories := repository.StartRepositories(database.MongoClient.Database("fibermongo"))
	services := service.StartServices(repositories)
	handlers := handlers.StartHandlers(services)

	s.mountRoutes(handlers)

	return s.app.Listen(":3000")
}

func (s *Server) mountRoutes(h *handlers.Handlers) {
	s.app.Get("/", h.HealthCheck.HealthCheck)
	routes.Users(s.app, h.Users)
}

func apiErrorHandler(c *fiber.Ctx, err error) error {
	apiErr, ok := err.(*schemas.ApiErr)
	if ok {
		return c.Status(apiErr.Code).JSON(fiber.Map{
			"error":  apiErr.Error(),
			"status": apiErr.Code,
		})
	}
	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
