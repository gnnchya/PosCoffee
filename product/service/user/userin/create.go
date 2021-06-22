package userin

import (
	"github.com/gnnchya/PosCoffee/product/domain"
)


type CreateInput struct {
	ID         		string   		`bson:"_id" json:"id"`
	Cart			domain.Cart   			`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		domain.GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"date_time" json:"date_time"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
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
		Code: 			input.Code,
		Err: 			input.Err,
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
