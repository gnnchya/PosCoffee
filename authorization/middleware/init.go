package middleware

import (
	"github.com/gnnchya/PosCoffee/authorize/config"
	"github.com/gnnchya/PosCoffee/authorize/service/grpcClient"
	"github.com/gnnchya/PosCoffee/authorize/service/roles"
)

type Service struct {
	conf         *config.Config
	grpcService  grpcClient.Service
	rolesService roles.Service
}

func New(conf *config.Config, grpcService grpcClient.Service, rolesService roles.Service) Service {
	return Service{conf, grpcService, rolesService}
}
