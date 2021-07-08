package wrapper

import (
	rolesService "github.com/gnnchya/PosCoffee/authorize/service/roles"
)

type wrapper struct {
	service rolesService.Service
}

func _(service rolesService.Service) rolesService.Service {
	return &wrapper{
		service: service,
	}
}
