package grpcClient

import (
	pb "github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf/cart"
)

type Service interface {
	GetMenu(input *pb.Request)(res *pb.Reply, err error)
}

