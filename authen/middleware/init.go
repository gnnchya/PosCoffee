package middleware

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	"github.com/gnnchya/PosCoffee/authen/service/user"
	"github.com/gnnchya/PosCoffee/authen/service/util"
)

type Service struct {
	Auth  authentication.Service
	Users user.Service
	Repo util.Repository
	Filter util.Filters
}

func New(authService authentication.Service, Users user.Service,repo util.Repository,filter util.Filters) Service {
	return Service{
		Auth:  authService,
		Users: Users,
		Repo: repo,
		Filter: filter,
	}
}
