package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type CreateInput struct {
	ID 				string  `json:"_id"`
	CustomerID 		string  `json:"customer_id"`
	Menu			[]domain.Menu 	`json:"menu"`
	Price           int64	`json:"price"`
	Code 			int 	`json:"code"`
	Err 			error	`json:"err"`
}

type CreateMenu struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
}

func (input *CreateInput)CreateInputToUserDomain() (user *domain.CreateStruct) {
	return &domain.CreateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Menu: input.Menu,
	}
}
