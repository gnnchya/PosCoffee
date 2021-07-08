package wrapper

import (
	permissionsService "github.com/gnnchya/PosCoffee/authorize/service/permissions"
)

type wrapper struct {
	service permissionsService.Service
}

func _(service permissionsService.Service) permissionsService.Service {
	return &wrapper{
		service: service,
	}
}
