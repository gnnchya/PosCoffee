package implement

import (
	"github.com/gnnchya/PosCoffee/product/service/grpcClient"
	"github.com/gnnchya/PosCoffee/product/service/user"
	"github.com/gnnchya/PosCoffee/product/service/util"
	"github.com/gnnchya/PosCoffee/product/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	repoMoney     util.RepositoryMoney
	mBroker   util.RepositoryMsgBroker
	client    grpcClient.Service
}

func New(validator validator.Validator, repo util.Repository,repoMoney util.RepositoryMoney, mBroker util.RepositoryMsgBroker, client grpcClient.Service) (service user.Service) {
	return &implementation{validator, repo, repoMoney, mBroker, client}
}
