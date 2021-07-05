package userin

import "github.com/gnnchya/PosCoffee/product/domain"

type ReadMoneyInput struct {
	ID string `json:"id"`
}


func (input *ReadMoneyInput)ReadMoneyInputToUserDomain() (user *domain.ReadMoneyStruct) {
	return &domain.ReadMoneyStruct{
		ID: input.ID,
	}
}

