package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/grpc/protobuf"
)

func (impl implementation) SendCart(ctx context.Context, request *protobuf.Request) (*protobuf.Reply, error){
	fmt.Println("here is the request from cart app:", request)
	output := &protobuf.Reply{
		Id : 123,
	}

	return output, nil
}