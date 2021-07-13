package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
)

func (impl *implementation) CheckPermission(ctx context.Context, roles []string, method string, path string)(bool, error){
	request := &protobuf.RequestPermission{
		Roles:      roles,
		Permission: method+path,
	}
	out, err := impl.client.CheckPermission(request)
	if err != nil{
		return false, err
	}
	return out.Allow, nil
}