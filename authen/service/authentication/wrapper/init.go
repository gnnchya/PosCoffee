package wrapper

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
)

type wrapper struct {
	service authentication.Service
}

func _(service authentication.Service) authentication.Service {
	return &wrapper{
		service: service,
	}
}
