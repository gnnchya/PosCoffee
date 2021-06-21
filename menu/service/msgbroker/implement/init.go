package implement

import (
	"github.com/gnnchya/PosCoffee/menu/service/msgbroker"
	userService "github.com/gnnchya/PosCoffee/menu/service/user"
	"github.com/gnnchya/PosCoffee/menu/service/util"
	)

type implementation struct{
	msgBroker util.RepositoryMsgBroker
	usrService userService.Service
}

func New(msgBroker util.RepositoryMsgBroker, usrService userService.Service) (service msgbroker.Service){
	impl := &implementation{
		msgBroker: msgBroker,
		usrService: usrService,
	}
	return impl
}