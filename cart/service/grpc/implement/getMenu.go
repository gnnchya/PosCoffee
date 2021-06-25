package implement

import (
	"context"
	pb "github.com/gnnchya/PosCoffee/cart/service/grpc/protobuf/cart"
)

func (impl implementation) GetMenu(ctx context.Context, request *pb.Request) (*pb.Reply, error) {
	input := request.Id
	out:=&pb.Reply{
		Id:  input,
		Arr: "hi from get menu",
	}
	return out, nil
}
