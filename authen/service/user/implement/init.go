package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/util"
)

type implementation struct {
	//validator       validator.Validator
	repo            util.RepositoryUsers
	//uuid            util.UUID
	filter          util.Filters
	mBroker         util.RepositoryMsgBroker
	cryptPassPhrase string
	//client          grpcClient.Service
}

func New(validator validator.Validator, repo util.RepositoryUsers, uuid util.UUID, filter util.Filters, mBroker util.RepositoryMsgBroker, cryptPassPhrase string, client grpcClient.Service) (service users.Service) {
	return &implementation{validator, repo, uuid, filter, mBroker, cryptPassPhrase, client}
}
