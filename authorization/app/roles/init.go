package roles

import (
	"github.com/gnnchya/PosCoffee/authorize/service/roles"
)

type Controller struct {
	service roles.Service
}

func New(service roles.Service) (ctrl *Controller) {
	return &Controller{service}
}