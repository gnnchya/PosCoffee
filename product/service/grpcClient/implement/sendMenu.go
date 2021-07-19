package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) SendMenu(input *protobuf.RequestMenu) (*protobuf.ReplyMenu, error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientMenu.SendMenu(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}
