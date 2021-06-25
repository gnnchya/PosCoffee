package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/grpc/protobuf"
)

func (impl implementation) SendCart(ctx context.Context, request *protobuf.Request2) (*protobuf.Reply2, error){
	fmt.Println("here is the request from cart app:", request)
	//TODO get change and return
	output := &protobuf.Reply2{
		Id : 123,
	}
	return output, nil
}