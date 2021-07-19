package middleware

import (
	"github.com/gnnchya/PosCoffee/product/service/user"
)

type Service struct {
	Users user.Service
}

func New(Users user.Service) Service {
	return Service{
		Users: Users,
	}
}
