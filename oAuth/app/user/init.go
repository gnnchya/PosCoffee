package user

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/grpcClient"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer"
)

type Controller struct {
	service consumer.Service
	grpcClient grpcClient.Service
}

func New(service consumerz.Service, grpcClient grpcClient.Service) (ctrl *Controller) {
	return &Controller{service, grpcClient}
}
