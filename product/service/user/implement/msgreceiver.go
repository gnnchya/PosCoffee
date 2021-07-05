package implement

import (
	"encoding/json"
	//"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) MsgReceiver(msg []byte) (err error) {
	msgInput := &userin.MsgBrokerCreate{}
	err = json.Unmarshal(msg, msgInput)
	if err != nil {
		return err
	}
	return
}
