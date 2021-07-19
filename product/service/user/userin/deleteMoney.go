package userin

import "github.com/gnnchya/PosCoffee/product/domain"

type DeleteMoneyInput struct {
	ID string `bson:"_id" json:"id"`
}



func (input *DeleteMoneyInput)DeleteMoneyInputToUserDomain() (user *domain.DeleteMoneyStruct) {
	return &domain.DeleteMoneyStruct{
		ID:             input.ID,
	}
}
