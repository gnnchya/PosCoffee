package consumer

import "github.com/gnnchya/PosCoffee/email/service/consumer"

type Controller struct {
	service consumer.Service
}

func New(service consumer.Service) (ctrl *Controller) {
	return &Controller{service}
}
