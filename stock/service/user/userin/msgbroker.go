package userin

import "github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"

type MsgBrokerCreate struct{
	Action 	msgbrokerin.ActionMsgBroker `json:"action"`
	ID         		string   	`bson:"_id" json:"id"`
	ItemName       	string   	`bson:"item_name" json:"item_name"`
	Category 		string  	`bson:"category" json:"category"`
	Amount			int64   	`bson:"amount" json:"amount"`
	Unit     		string   	`bson:"unit" json:"unit"`
	CostPerUnit		int64      	`bson:"cost_per_unit" json:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date" json:"exp_date"`
	ImportDate      int64   	`bson:"import_date" json:"import_date"`
	Supplier 		string 		`bson:"supplier" json:"supplier"`
	TotalCost		int64      	`bson:"total_cost" json:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount" json:"total_amount"`

}

func (msg MsgBrokerCreate) ToCreateInput()(createInput *CreateInput){
	createInput = &CreateInput{
		ID:             msg.ID,
		ItemName:       msg.ItemName,
		Category:     	msg.Category,
		Amount: 		msg.Amount,
		Unit:         	msg.Unit,
		CostPerUnit:    msg.CostPerUnit,
		EXPDate:        msg.EXPDate,
		ImportDate:     msg.ImportDate,
		Supplier:       msg.Supplier,
		TotalCost:      msg.TotalCost,
		TotalAmount:    msg.TotalAmount,

	}
	return createInput
}

type MsgBrokerUpdate struct {
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	Amount	int64   `bson:"amount" json:"amount"`
	ID 		string	`bson:"_id" json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

func (msg MsgBrokerCreate) ToUpdateInput()(input *UpdateInput) {
	input = &UpdateInput{
		ID:         msg.ID,
		Amount:     msg.Amount,
	}
	return input
}

type MsgBrokerDelete struct {
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID 		string `json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

func (msg MsgBrokerCreate) ToDeleteInput()(input *DeleteInput) {
	input = &DeleteInput{
		ID: 		msg.ID,
	}
	return input
}