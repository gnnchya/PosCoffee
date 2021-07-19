package grpcClient

import "github.com/gnnchya/PosCoffee/cart/service/grpcClient/protobuf"

type Service interface {
	SendCart(input *protobuf.Request2) (*protobuf.Reply2, error)
	Middleware(input *protobuf.RequestMiddleware) (*protobuf.ReplyMiddleware, error)
}

