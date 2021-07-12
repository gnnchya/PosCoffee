package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf"
)

func (impl implementation) CheckPermission(ctx context.Context, input *protobuf.RequestPermission) (output *protobuf.ReplyPermission, err error) {
	role, err := impl.roleService.CheckPermission(ctx, input.Roles, input.Permission)
	if err != nil{
		return nil, err
	}
	out := &protobuf.ReplyPermission{Allow: role}

	return out, nil
}

