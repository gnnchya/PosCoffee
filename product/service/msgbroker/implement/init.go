package implement

import (
	"github.com/gnnchya/PosCoffee/product/service/msgbroker"
	userService "github.com/gnnchya/PosCoffee/product/service/user"
	"github.com/gnnchya/PosCoffee/product/service/util"

	//"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user"
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