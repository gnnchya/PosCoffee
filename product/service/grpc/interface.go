package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/grpc/protobuf"
)

type Service interface {
	SendCart(context.Context, *protobuf.Request) (*protobuf.Reply, error)
}