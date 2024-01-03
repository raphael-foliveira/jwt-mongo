package service

import "github.com/raphael-foliveira/fiber-mongo/internal/api/repository"

type Services struct {
	Users Users
	Jwt   JwtService
}

func StartServices(repositories *repository.Repositories) *Services {
	return &Services{
		Users: NewUsersService(repositories.Users, NewJwtService()),
	}
}
