package out

import "github.com/gnnchya/PosCoffee/authen/service/msgbroker/msgbrokerin"

type MsgBroker struct {
	Action       msgbrokerin.ActionMsgBroker `json:"action"`
	Channel      string                      `json:"channel"`
	ID           string                      `json:"id"`
	Email        string                      `json:"email"`
	MobileNumber string                      `json:"mobile_number"`
}
