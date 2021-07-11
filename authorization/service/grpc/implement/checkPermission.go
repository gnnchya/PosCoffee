package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (impl *implementation) CheckPermission(data *protobuf.RequestPermission) (res *protobuf.ReplyPermission, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = impl.clientPermission.CheckPermission(ctx, data)

	if status.Code(err) != codes.OK {
		return nil, err
	}
	return res, nil
}

