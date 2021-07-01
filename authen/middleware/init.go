package middleware

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	"github.com/gnnchya/PosCoffee/authen/service/user"
)

type Service struct {
	Auth  authentication.Service
	Users user.Service
}

func New(authService authentication.Service, Users user.Service) Service {
	return Service{
		Auth:  authService,
		Users: Users,
	}
}
