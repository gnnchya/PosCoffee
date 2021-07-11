package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) CheckPermission(input *protobuf.RequestPermission) (*protobuf.ReplyPermission, error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.client.CheckPermission(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}