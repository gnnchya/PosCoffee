package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

//type Cart struct{
//	Menu		Menu `bson:"menu" json:"menu"`
//	Amount 		int64   `bson:"amount" json:"amount"`
//	Option 		string   `bson:"option" json:"option"`
//}
//
//type Menu struct{
//	ID         		string   `bson:"_id" json:"id"`
//	Category       	string   `bson:"category" json:"category"`
//	Name 			string   `bson:"name" json:"name" validate:"required"`
//	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
//	Price      		int64    `bson:"price" json:"price"`
//	Available 		bool	 `bson:"available" json:"available"`
//	Code int `json:"code"`
//	Err error `json:"err"`
//}

type CreateInput struct {
	ID         		string   `bson:"_id" json:"id"`
	CustomerID 		string   `bson:"customer_id" json:"customer_id"`
	Cart			[]domain.Cart		`bson:"cart" json:"cart"`
	Purchase     		bool   `bson:"status" json:"status"`
	Price  			int64    `bson:"price" json:"price"`
	TypeOfOrder 	string `bson:"type_of_order" json:"type_of_order"`
	Destination    	string     `bson:"destination" json:"destination"`
	Time			int64 	`bson:"time" json:"time"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func (input *CreateInput)CreateInputToUserDomain() (user *domain.CreateStruct) {
	return &domain.CreateStruct{
		ID:             input.ID,
		CustomerID: input.CustomerID,
		Cart: input.Cart,
		Purchase: input.Purchase,
		Price: input.Price,
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
