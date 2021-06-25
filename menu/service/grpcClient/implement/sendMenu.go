package implement

import (
	"context"
	pb "github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf/cart"
	"time"
)

func (impl *implementation) GetMenu(input *pb.Request) (res *pb.Reply, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.client.GetMenu(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}
