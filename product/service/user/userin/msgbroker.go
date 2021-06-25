package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
)

type MsgBrokerCreate struct{
	Action msgbrokerin.ActionMsgBroker `json:"action"`
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

func (msg MsgBrokerCreate) ToCreateInput()(input *domain.CreateStockStruct){
	input = &domain.CreateStockStruct{
		ID:             input.ID,
		ItemName:       input.ItemName,
		Category:     	input.Category,
		Amount: 		input.Amount,
		Unit:         	input.Unit,
		CostPerUnit:    input.CostPerUnit,
		EXPDate:        input.EXPDate,
		ImportDate:     input.ImportDate,
		Supplier:       input.Supplier,
		TotalCost:      input.TotalCost,
		TotalAmount:    input.TotalAmount,
	}
	return input
}

type MsgBrokerUpdate struct {
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID      string  `bson:"_id" json:"id"`
	Amount	int64   `bson:"amount" json:"amount"`
	Code 	int 	`json:"code"`
	Err 	error   `json:"err"`
}

func (msg MsgBrokerCreate) ToUpdateInput()(input *domain.UpdateStockStruct) {
	input = &domain.UpdateStockStruct{
		ID:        	input.ID,
		Amount:   	input.Amount,
		Code: 		input.Code,
		Err: 		input.Err,
	}
	return input
}

type MsgBrokerDelete struct {
	Action msgbrokerin.ActionMsgBroker `json:"action"`
	ID string `bson:"_id" json:"id"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func (msg MsgBrokerCreate) ToDeleteInput()(input *domain.DeleteStockStruct) {
	input = &domain.DeleteStockStruct{
		ID:   input.ID,
		Code: input.Code,
		Err:  input.Err,
	}
	return input
}