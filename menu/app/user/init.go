package user

import (
	"github.com/gnnchya/PosCoffee/menu/service/user"
	grpcClient "github.com/gnnchya/PosCoffee/menu/service/grpcClient"
)

type Controller struct {
	service user.Service
	grpcClient  grpcClient.Service
}

func New(service user.Service, grpcClient  grpcClient.Service) (ctrl *Controller) {
	return &Controller{service, grpcClient}
}
