package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) ReadStock(input *protobuf.RequestRead) (*protobuf.ReplyRead, error) {
	fmt.Println("send read to stock", input)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientRead.ReadStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}
