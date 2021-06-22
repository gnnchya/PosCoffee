package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type UpdateInput struct {
	ID         		string   `bson:"_id" json:"id"`
	CustomerID 		string   `bson:"customer_id" json:"customer_id"`
	Cart  			[]struct{
		ID         		string   `bson:"_id" json:"id"`
		Category       	string   `bson:"category" json:"category"`
		Name 			string   `bson:"name" json:"name" validate:"required"`
		Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
		Price      		int64    `bson:"price" json:"price"`
		Available 		bool	 `bson:"available" json:"available"`
		Code int `json:"code"`
		Err error `json:"err"`
	}   `bson:"cart" json:"cart"`
	Status     		string   `bson:"status" json:"status"`
	Price  			int64    `bson:"price" json:"price"`
	TypeOfOrder 	string `bson:"type_of_order" json:"type_of_order"`
	Destination    	string     `bson:"destination" json:"destination"`
	Time			int64 	`bson:"time" json:"time"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func UpdateInputToUserDomain(input *UpdateInput) (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Cart: input.Cart,
		Status: input.Status,
		Price: input.Price,
		TypeOfOrder: input.TypeOfOrder,
		Destination: input.Destination,
		Code: input.Code,
		Err: input.Err,
	}
}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Cart: input.Cart,
		Status: input.Status,
		Price: input.Price,
		TypeOfOrder: input.TypeOfOrder,
		Destination: input.Destination,
		Code: input.Code,
		Err: input.Err,
	}
}