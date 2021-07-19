package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) SendIngredients(input *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.client.SendIngredients(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}