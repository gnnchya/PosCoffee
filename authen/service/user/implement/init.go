package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/user"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"github.com/gnnchya/PosCoffee/authen/service/validator"
)

type implementation struct {
	validator validator.Validator
	elasRepo  util.RepositoryElas
	redisRepo util.RepositoryRedis
}

func New(validator validator.Validator, elasRepo util.RepositoryElas, redisRepo util.RepositoryRedis) (service user.Service) {
	return &implementation{validator, elasRepo, redisRepo}
}
