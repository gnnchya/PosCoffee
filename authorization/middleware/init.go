package middleware

import (
	"github.com/gnnchya/PosCoffee/authorize/config"
	"github.com/gnnchya/PosCoffee/authorize/service/grpc"
	"github.com/gnnchya/PosCoffee/authorize/service/roles"
)

type Service struct {
	conf         *config.Config
	grpcService  grpc.Service
	rolesService roles.Service
}

func New(conf *config.Config, grpcService grpc.Service, rolesService roles.Service) Service {
	return Service{conf, grpcService, rolesService}
}
