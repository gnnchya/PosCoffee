package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/grpc/protobuf"
)

type Service interface {
	Middleware(ctx context.Context, input *protobuf.RequestMiddleware) (*protobuf.ReplyMiddleware, error)
}

