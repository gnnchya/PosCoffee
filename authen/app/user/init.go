package user

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient"
	"github.com/gnnchya/PosCoffee/authen/service/user"
)

type Controller struct {
	service user.Service
	authService authentication.Service
	grpcClient  grpcClient.Service
}

func New(service user.Service, authService authentication.Service, grpcClient grpcClient.Service) (ctrl *Controller) {
	return &Controller{service, authService, grpcClient}
}
