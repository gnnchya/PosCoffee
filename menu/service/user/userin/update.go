package userin

import (
	"github.com/gnnchya/PosCoffee/menu/domain"
)

type UpdateInput struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]domain.Ingredient `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.UpdateStruct) {
	return &domain.UpdateStruct{
		ID:             input.ID,
		Category: input.Category,
		Name:           input.Name,
		Ingredient: input.Ingredient,
		Price: input.Price,
		Available: input.Available,
	}
}