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
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	app         *fiber.App
	mongoClient *mongo.Client
}

func NewServer(c *mongo.Client) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: apiErrorHandler,
	})
	return &Server{app, c}
}

func (s *Server) Start() error {
	s.app.Use(cors.New())
	s.app.Use(logger.New())
	s.app.Get("/", handlers.HealthCheck)
	s.Bootstrap()
	return s.app.Listen(":3000")
}

func (s *Server) Bootstrap() {
	usersRepo := repository.NewUsers(s.mongoClient)
	usersService := service.NewUsers(usersRepo)
	usersHandler := handlers.NewUsers(usersService)
	routes.Users(s.app, usersHandler)
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
