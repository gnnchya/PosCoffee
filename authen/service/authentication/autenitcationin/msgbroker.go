package autenitcationin

import "github.com/gnnchya/PosCoffee/authen/service/msgbroker/msgbrokerin"

type MsgBrokerForgot struct {
	Action         msgbrokerin.ActionMsgBroker `json:"action"`
	ID             string                      `json:"id"`
	SignUpChannel  string                      `json:"sign_up_channel"`
	Email          string                      `json:"email"`
	MobileNumber   string                      `json:"mobile_number" validate:"required"`
	IdentifyType   string                      `json:"identify_type" validate:"required"`
	IdentifyNumber string                      `json:"identify_number" validate:"required"`
	Password       string                      `json:"password" validate:"required"`
	Passcode       string                      `json:"passcode"`
}

type MsgBrokerVerify struct {
	Action  msgbrokerin.ActionMsgBroker `json:"action"`
	Code    string                      `json:"code" validate:"required"`
	RefCode string                      `json:"ref_code" validate:"required"`
	Contact string                      `json:"contact" validate:"required"`
}
