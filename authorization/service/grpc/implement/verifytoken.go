package implement

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf/authen"
)

func (impl *implementation) VerifyToken(data *pb.TokenRequest) (res *pb.TokenReply, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = impl.client.VerifyToken(ctx, data)

	if status.Code(err) != codes.OK {
		return nil, err
	}
	return res, nil
}
