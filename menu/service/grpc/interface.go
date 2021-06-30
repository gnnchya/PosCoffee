package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/service/grpc/protobuf"
)

type Service interface {
	SendMenu(ctx context.Context, request *protobuf.RequestMenu) (*protobuf.ReplyMenu, error)
}
