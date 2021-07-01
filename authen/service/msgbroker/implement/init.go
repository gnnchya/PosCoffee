package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/msgbroker"
	userService "github.com/gnnchya/PosCoffee/authen/service/user"
	"github.com/gnnchya/PosCoffee/authen/service/util"

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