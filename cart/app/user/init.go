package user

import (
	"github.com/gnnchya/PosCoffee/cart/service/user"
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient"
)

type Controller struct {
	service user.Service
	grpcClient grpcClient.Service
}

func New(service user.Service, grpcClient grpcClient.Service) (ctrl *Controller) {
	return &Controller{service, grpcClient}
}
