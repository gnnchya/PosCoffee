package token

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/token"
)

type Controller struct {
	service token.Service
}

func New(service token.Service) (ctrl *Controller) {
	return &Controller{service}
}
