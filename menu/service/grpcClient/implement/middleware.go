package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) Middleware(input *protobuf.RequestMiddleware) (*protobuf.ReplyMiddleware, error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientMiddleware.Middleware(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}