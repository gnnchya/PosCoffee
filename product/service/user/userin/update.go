package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
)


type UpdateInput struct {
	ID         		string   		`bson:"_id" json:"id"`
	Finished		bool     		`bson:"finished" json:"finished"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		domain.GeoJson  `bson:"destination" json:"destination"`
}


func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.UpdateOrderStruct) {
	return &domain.UpdateOrderStruct{
		ID:             input.ID,
		Finished: 		input.Finished,
		TypeOfOrder: 	input.TypeOfOrder,
		Destination: 	input.Destination,
	}
}