package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient/protobuf"
)

func (impl *implementation) Finish(ctx context.Context, id string) (ID string, err error) {
	input, err := impl.repo.Read(ctx, id)
	menu := toMenuArr(input.Menu)
	inputProtobuf := &protobuf.Request{
		Cart: &protobuf.Cart{
			Id:         id,
			CustomerId: input.CustomerID,
			Menu:      menu,
			Price:      input.TotalPrice,
		},
	}
	response, err:= impl.client.SendCart(inputProtobuf)
	fmt.Println("response", response)
	return input.ID, nil
}

func toMenuArr(menu []domain.Menu)(a []*protobuf.Menu){
	for _, key := range menu{
		var ingre []*protobuf.Ingredient
		for _, key2 := range key.Ingredient{
			toProtoIngre := &protobuf.Ingredient{
				IngredientName: key2.IngredientName,
				Amount:         key2.Amount,
			}
			ingre = append(ingre, toProtoIngre)
		}
		toProto := &protobuf.Menu{
			Id:         key.ID,
			Category:   key.Category,
			Name:       key.Name,
			Ingredient: ingre,
			Price:      key.Price,
			Available:  key.Available,
		}
		a = append(a, toProto)
	}
	return a
}