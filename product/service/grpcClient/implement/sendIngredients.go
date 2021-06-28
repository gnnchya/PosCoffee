package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) SendIngredients(input *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error){
	fmt.Println("input to stock", input)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.client.SendIngredients(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}