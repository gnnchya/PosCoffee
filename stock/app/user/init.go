package user

import (
	"github.com/gnnchya/PosCoffee/stock/service/user"
)

type Controller struct {
	service user.Service
}

func New(service user.Service) (ctrl *Controller) {
	return &Controller{service}
}
