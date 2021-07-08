package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/user"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"github.com/gnnchya/PosCoffee/authen/service/validator"
)

type implementation struct {
	validator       validator.Validator
	repo            util.RepositoryUsers
	//uuid            util.UUID
	filter          util.Filters
	mBroker         util.RepositoryMsgBroker
	cryptPassPhrase string
	//client          grpcClient.Service
}

func New(validator validator.Validator, repo util.RepositoryUsers, filter util.Filters, mBroker util.RepositoryMsgBroker, cryptPassPhrase string) (service user.Service) {
	return &implementation{validator, repo, filter, mBroker, cryptPassPhrase}
}
