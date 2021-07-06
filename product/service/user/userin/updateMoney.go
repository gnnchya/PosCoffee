package userin

import "github.com/gnnchya/PosCoffee/product/domain"

type UpdateMoneyInput struct {
	ID 		string	`bson:"_id" json:"id"`
	Value   int64   `bson:"value" json:"value"`
	Amount 	int64  	`bson:"amount" json:"amount"`
}

func (input *UpdateMoneyInput)UpdateStockInputToUserDomain() (user *domain.UpdateMoneyStruct) {
	return &domain.UpdateMoneyStruct{
		ID:     input.ID,
		Value:  input.Value,
		Amount: input.Amount,
	}
}
