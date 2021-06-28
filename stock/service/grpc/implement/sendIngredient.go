package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"
)

func(impl implementation) SendIngredients(ctx context.Context, input *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error){
	fmt.Println("from product", input)
	out := &protobuf.ReplyFromStock{
		Stock:         false,
		CalculateCost: nil,
		Err:           "test",
	}
	return out, nil
}