package user

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	"github.com/gnnchya/PosCoffee/authen/service/user"
)

type Controller struct {
	service     user.Service
	authService authentication.Service
}

func New(service user.Service, authService authentication.Service) (ctrl *Controller) {
	return &Controller{service, authService}
}