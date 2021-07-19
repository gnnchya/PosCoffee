package implement

import (
	"github.com/gnnchya/PosCoffee/menu/service/grpcClient"
	"github.com/gnnchya/PosCoffee/menu/service/user"
	"github.com/gnnchya/PosCoffee/menu/service/util"
	"github.com/gnnchya/PosCoffee/menu/service/validator"
)

type implementation struct {
	validator validator.Validator
	elasRepo  util.RepositoryElas
	redisRepo util.RepositoryRedis
	client grpcClient.Service
}

func New(validator validator.Validator, elasRepo util.RepositoryElas, redisRepo util.RepositoryRedis, client grpcClient.Service) (service user.Service) {
	return &implementation{validator, elasRepo, redisRepo, client}
}
