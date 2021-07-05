package userin

import "github.com/gnnchya/PosCoffee/product/domain"

type CreateMoneyInput struct {
	ID 			string		`bson:"_id" json:"id"`
	Value   	int64   	`bson:"value" json:"value"`
	Amount 		int64  		`bson:"amount" json:"amount"`
	Currency	string   	`bson:"currency" json:"currency"`
}

func (input *CreateMoneyInput)CreateMoneyInputToUserDomain() (user *domain.CreateMoneyStruct) {
	return &domain.CreateMoneyStruct{
		ID:       input.ID,
		Value:    input.Value,
		Amount:   input.Amount,
		Currency: input.Currency,
	}
}
