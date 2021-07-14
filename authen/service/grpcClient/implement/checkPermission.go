package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
	"time"
)

func (impl *implementation) CheckPermission(input *protobuf.RequestPermission) (*protobuf.ReplyPermission, error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	fmt.Println("in check permission")
	r, err := impl.client.CheckPermission(ctx, input)
	fmt.Println("err", err)
	if err != nil {
		return r, err
	}
	return r, nil
}