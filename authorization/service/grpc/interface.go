package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf"
)

//go:generate mockery --name=Service
type Service interface {
	//VerifyToken(data *pb.TokenRequest) (res *pb.TokenReply, err error)
	CheckPermission(context.Context, *protobuf.RequestPermission) (*protobuf.ReplyPermission, error)
}
