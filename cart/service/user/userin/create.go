package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)


type CreateInput struct {
	ID         		string   `bson:"_id" json:"id"`
	CustomerName    string   `bson:"customer_name" json:"customer_name" validate:"required"`
	CustomerID 		string   `bson:"customer_id" json:"customer_id"`
	Cart  			string   `bson:"actual_lastname" json:"actual_lastname"`
	Status     		string   `bson:"status" json:"status"`
	Price  			int64    `bson:"birth_date" json:"birth_date"`
	Branch     		string      `bson:"branch" json:"branch"`
	TypeOfOrder 	string `bson:"type_of_order" json:"type_of_order"`
	Destination    	string     `bson:"destination" json:"destination"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func (input *CreateInput)CreateInputToUserDomain() (user *domain.InsertQ) {
	return &domain.InsertQ{
		ID:             input.ID,
		CustomerName: input.CustomerName,
		CustomerID: input.CustomerID,
		Cart: input.Cart,
		Status: input.Status,
		Price: input.Price,
		Branch: input.Branch,
		TypeOfOrder: input.TypeOfOrder,
		Destination: input.Destination,
		Code: input.Code,
		Err: input.Err,
	}
}

//func ToDomain(input *CreateInput) (user *domain.InsertQ) {
//
//	return &domain.InsertQ{
//		ID:         input.ID,
//		Name:       input.Name,
//		ActualName: input.ActualName,
//		Gender:     input.Gender,
//		BirthDate:  input.BirthDate,
//		Height:     input.Height,
//		SuperPower: input.SuperPower,
//		Alive:      input.Alive,
//	}
//}
