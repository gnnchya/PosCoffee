package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/service/grpc/protobuf"
	//repo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
)

func (impl implementation) SendMenu(ctx context.Context, request *protobuf.RequestMenu) (*protobuf.ReplyMenu, error){
	//TODO read all
	fmt.Println("enter send menu grpc", request)
	input, err := impl.userService.Report(ctx)
	if err != nil{
		return nil, err
	}
	var output []*protobuf.Menu
	for _, k := range input{
		var ingredient []*protobuf.MenuIngredient
		for _, in := range k.Ingredient{
			ingre := &protobuf.MenuIngredient{
				IngredientName: in.IngredientName,
				Amount:         in.Amount,
			}
			ingredient = append(ingredient, ingre)
		}
		out := &protobuf.Menu{
			ID:         k.ID,
			Category:   k.Category,
			Name:       k.Name,
			Ingredient: ingredient,
			Price:      k.Price,
			Available:  k.Available,
		}
		output = append(output, out)
	}

	result := &protobuf.ReplyMenu{Menu: output}

	return result, nil
}
