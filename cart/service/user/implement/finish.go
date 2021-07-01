package implement

import (
	"context"
	"errors"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (impl *implementation) Finish(ctx context.Context, id string, finishInput *userin.FinishInput) (message interface{}, res []string,err error) {
	input, err := impl.repo.Read(ctx, id)
	if err != nil{
		return "",res, err
	}
	menu := toMenuArr(input.Menu)
	inputProtobuf := &protobuf.Request2{
		Cart: &protobuf.Cart2{
			Id:         id,
			CustomerId: input.CustomerID,
			Menu:      menu,
			Price:      input.TotalPrice,
		},
		Paid: finishInput.Paid,
		PaymentMethod: finishInput.PaymentMethod,
		Type: finishInput.TypeOfOrder,
		Geo: &protobuf.GeoJson{
			Type: "point",
			Long: finishInput.Longitude,
			Lat:  finishInput.Latitude,
		},
	}
	fmt.Println("inputProto", inputProtobuf)
	response, err := impl.client.SendCart(inputProtobuf)
	fmt.Println("response from product", response)
	fmt.Println("error", err)
	if err != nil {
		return "", res, err
	}
	if response.Stock == true {
		return response.Changes, response.Bill, nil
	}else{
		return input.ID, res, errors.New(response.Err)
	}
}

func toMenuArr(menu []domain.Menu)(a []*protobuf.Menu2){
	for _, key := range menu{
		var ingre []*protobuf.Ingredient2
		for _, key2 := range key.Ingredient{
			toProtoIngre := &protobuf.Ingredient2{
				IngredientName: key2.ItemName,
				Amount:         key2.Amount,
			}
			ingre = append(ingre, toProtoIngre)
		}
		toProto := &protobuf.Menu2{
			Id:         key.ID,
			Category:   key.Category,
			Name:       key.Name,
			Ingredient: ingre,
			Price:      key.Price,
			Available:  key.Available,
			Amount: key.Amount,
			Option: key.Option,
		}
		a = append(a, toProto)
	}
	return a
}