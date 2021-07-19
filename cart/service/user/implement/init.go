package implement

import (
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	"github.com/gnnchya/PosCoffee/cart/service/user"
	"github.com/gnnchya/PosCoffee/cart/service/util"
	"github.com/gnnchya/PosCoffee/cart/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	client grpcClient.Service
}

func New(validator validator.Validator, repo util.Repository, grpcClient grpcClient.Service) (service user.Service) {
	return &implementation{validator, repo, grpcClient}
}
