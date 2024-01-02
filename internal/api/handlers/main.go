package handlers

import "github.com/raphael-foliveira/fiber-mongo/internal/api/service"

type Handlers struct {
	Users       Users
	HealthCheck *HealthCheck
}

func StartHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Users:       NewUsersHandler(services.Users),
		HealthCheck: NewHealthCheckHandler(),
	}
}
