package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) SendCart(input *protobuf.Request2) (*protobuf.Reply2, error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.client.SendCart(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}