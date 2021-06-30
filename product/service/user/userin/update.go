package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/product/domain"
)


type UpdateInput struct {
	ID         		string   		`bson:"_id" json:"id"`
	Cart			domain.Cart		`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		domain.GeoJson  `bson:"destination" json:"destination"`
	Time			int64      		`bson:"date_time" json:"date_time"`
}

//func UpdateInputToUserDomain(input *UpdateInput) (user *domain.UpdateOrderStruct) {
//	return &domain.UpdateOrderStruct{
//		ID:             input.ID,
//		Cart:
//		Code:			input.Code,
//		Err: input.Err,
//	}
//}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.UpdateOrderStruct) {
	return &domain.UpdateOrderStruct{
		ID:             input.ID,
		Cart: 			input.Cart,
		Finished: 		input.Finished,
		Price: 			input.Price,
		TypeOfOrder: 	input.TypeOfOrder,
		Destination: 	input.Destination,
		Time:			input.Time,
	}
}