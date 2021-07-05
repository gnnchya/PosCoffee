package user

import (
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	"github.com/gnnchya/PosCoffee/cart/service/user"
)

type Controller struct {
	service user.Service
	grpcClient grpcClient.Service
}

func New(service user.Service, grpcClient grpcClient.Service) (ctrl *Controller) {
	return &Controller{service, grpcClient}
}
