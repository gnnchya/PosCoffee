package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/totalcost"
)


type CreateInput struct {
	ID         		string   		`bson:"_id" json:"id"`
	Cart			domain.Cart   	`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	PaymentMethod	string			`bson:"payment_method" json:"payment_method"`
	Destination		domain.GeoJson  `bson:"destination" json:"destination"`
	Time			int64      		`bson:"date_time" json:"date_time"`
}

func (input *CreateInput)CreateInputToUserDomain() (user *domain.CreateOrderStruct) {
	return &domain.CreateOrderStruct{
		ID:             input.ID,
		Cart: 			input.Cart,
		Finished: 		input.Finished,
		Price: 			input.Price,
		TypeOfOrder: 	input.TypeOfOrder,
		Destination: 	input.Destination,
		Time:			input.Time,
		TotalCost:		totalCost.CalculateTotalCost,
	}
}
