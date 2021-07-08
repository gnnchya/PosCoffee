package permissions

import (
	"github.com/gnnchya/PosCoffee/authorize/service/permissions"
)

type Controller struct {
	service permissions.Service
}

func New(service permissions.Service) (ctrl *Controller) {
	return &Controller{service}
}
