package userin

import "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"

type MsgBrokerCreate struct{
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID         string   `json:"id"`
	Name       string   `son:"name" validate:"required"`
	ActualName string   `json:"actual_name" validate:"required"`
	ActualLastName string `json:"actual_last_name"`
	Gender     string   `json:"gender"`
	BirthDate  int64    `json:"birth_date"`
	Height     int      `json:"height" validate:"gte=0"`
	SuperPower []string `json:"super_power"`
	Alive      bool     `json:"alive"`
	Universe       string   `json:"universe"`
	Movies         []string `json:"movies"`
	Enemies        []string `json:"enemies"`
	FamilyMember   []string `json:"family_member"`
	About          string   `json:"about"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func (msg MsgBrokerCreate) ToCreateInput()(createInput *CreateInput){
	createInput = &CreateInput{
		ID:         msg.ID,
		Name:       msg.Name,
		ActualName: msg.ActualName,
		ActualLastName: msg.ActualLastName,
		Gender:     msg.Gender,
		BirthDate:  msg.BirthDate,
		Height:     msg.Height,
		SuperPower: msg.SuperPower,
		Alive:      msg.Alive,
		Universe: msg.Universe,
		Movies: msg.Movies,
		Enemies: msg.Enemies,
		FamilyMember: msg.FamilyMember,
		About: msg.About,
		Code: msg.Code,
		Err: msg.Err,
	}
	return createInput
}

type MsgBrokerUpdate struct {
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID         string   `json:"id"`
	Name       string   `son:"name" validate:"required"`
	ActualName string   `json:"actual_name" validate:"required"`
	ActualLastName string `json:"actual_last_name"`
	Gender     string   `json:"gender"`
	BirthDate  int64    `json:"birth_date"`
	Height     int      `json:"height" validate:"gte=0"`
	SuperPower []string `json:"super_power"`
	Alive      bool     `json:"alive"`
	Universe       string   `json:"universe"`
	Movies         []string `json:"movies"`
	Enemies        []string `json:"enemies"`
	FamilyMember   []string `json:"family_member"`
	About          string   `json:"about"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func (msg MsgBrokerCreate) ToUpdateInput()(input *UpdateInput) {
	input = &UpdateInput{
		ID:         msg.ID,
		Name:       msg.Name,
		ActualName: msg.ActualName,
		ActualLastName: msg.ActualLastName,
		Gender:     msg.Gender,
		BirthDate:  msg.BirthDate,
		Height:     msg.Height,
		SuperPower: msg.SuperPower,
		Alive:      msg.Alive,
		Universe: msg.Universe,
		Movies: msg.Movies,
		Enemies: msg.Enemies,
		FamilyMember: msg.FamilyMember,
		About: msg.About,
		Err: msg.Err,

	}
	return input
}

type MsgBrokerDelete struct {
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID string `json:"id"`
}

func (msg MsgBrokerCreate) ToDeleteInput()(input *DeleteInput) {
	input = &DeleteInput{
		ID: msg.ID,
	}
	return input
}