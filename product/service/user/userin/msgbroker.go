package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
)

type MsgBrokerCreate struct{
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID         		string   		`bson:"_id" json:"id"`
	Cart			domain.Cart   	`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		domain.GeoJson  `bson:"destination" json:"destination"`
	Time			int64      		`bson:"date_time" json:"date_time"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
}

func (msg MsgBrokerCreate) ToCreateInput()(input *CreateInput){
	input = &CreateInput{
		ID:             input.ID,
		Cart: 			input.Cart,
		Finished: 		input.Finished,
		Price: 			input.Price,
		TypeOfOrder: 	input.TypeOfOrder,
		Destination: 	input.Destination,
		Time:			input.Time,
		Code: 			input.Code,
		Err: 			input.Err,
	}
	return input
}

type MsgBrokerUpdate struct {
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID         		string   		`bson:"_id" json:"id"`
	Cart			domain.Cart   	`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		domain.GeoJson  `bson:"destination" json:"destination"`
	Time			int64      		`bson:"date_time" json:"date_time"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
}

func (msg MsgBrokerCreate) ToUpdateInput()(input *UpdateInput) {
	input = &UpdateInput{
		ID:             input.ID,
		Cart: 			input.Cart,
		Finished: 		input.Finished,
		Price: 			input.Price,
		TypeOfOrder: 	input.TypeOfOrder,
		Destination: 	input.Destination,
		Time:			input.Time,
		Code: 			input.Code,
		Err: 			input.Err,
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