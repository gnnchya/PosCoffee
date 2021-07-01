package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
)

func (impl implementation) newHandler(topic msgbrokerin.TopicMsgBroker) (handler common.Handler) {
	return func(ctx context.Context, msg []byte) {
		var err error
		switch topic {
		case msgbrokerin.TopicCreate:
			fmt.Println("enter newHandler responseCreate topic")
			err = impl.usrService.MsgReceiver(ctx, msg)
		case msgbrokerin.TopicUpdate:
			fmt.Println("enter newHandler responseUpdate topic")
			err = impl.usrService.MsgReceiver(ctx, msg)
		case msgbrokerin.TopicDelete:
			fmt.Println("enter newHandler responseDel topic")
			err = impl.usrService.MsgReceiver(ctx, msg)
		default:
			fmt.Println(string(msg), "with default")
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}