package grpcClient

import "github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf"

type Service interface {
	Middleware(input *protobuf.RequestMiddleware) (*protobuf.ReplyMiddleware, error)
}

