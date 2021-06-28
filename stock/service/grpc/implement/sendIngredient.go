package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"

)

func(impl implementation) SendIngredients(ctx context.Context, input *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error){
	fmt.Println("from product", input)
	state, expenses, err := impl.userService.CheckStock(ctx, input.Ingredient)
	var expense []*protobuf.CalculateCost
	for _, v := range expenses{
		temp := &protobuf.CalculateCost{
			ItemName:    v.ItemName,
			CostPerUnit: v.CostPerUnit,
		}
		expense = append(expense, temp)
	}
	out := &protobuf.ReplyFromStock{
		Stock:         state,
		CalculateCost: expense,
		Err:           err,
	}
	return out, nil
}