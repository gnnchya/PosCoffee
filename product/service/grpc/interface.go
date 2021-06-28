package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/grpc/protobuf"
)

type Service interface {
	SendCart(ctx context.Context, request *protobuf.Request2) (*protobuf.Reply2, error)
}

