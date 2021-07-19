package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf"
)

func (impl implementation) CheckPermission(ctx context.Context, input *protobuf.RequestPermission) (output *protobuf.ReplyPermission, err error) {
	fmt.Println("checkPermission gRPC", input)
	role, err := impl.roleService.CheckPermission(ctx, input.Roles, input.Permission)
	if err != nil{
		return nil, err
	}
	out := &protobuf.ReplyPermission{Allow: role}
	fmt.Println("out", out)
	return out, nil
}

