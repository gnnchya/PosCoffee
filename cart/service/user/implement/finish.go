package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient/protobuf"
)

func (impl *implementation) Finish(ctx context.Context, id string) (ID string, err error) {
	input, err := impl.repo.Read(ctx, id)
	fmt.Println("read input", input)
	menu := toMenuArr(input.Menu)
	inputProtobuf := &protobuf.Request2{
		Cart: &protobuf.Cart2{
			Id:         id,
			CustomerId: input.CustomerID,
			Menu:      menu,
			Price:      input.TotalPrice,
		},
	}
	fmt.Println("inputProto", inputProtobuf)
	response, err:= impl.client.SendCart(inputProtobuf)
	fmt.Println("response", response)
	return input.ID, nil
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