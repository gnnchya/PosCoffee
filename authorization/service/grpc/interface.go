package grpc

import (
	pb "github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf/authen"
)

//go:generate mockery --name=Service
type Service interface {
	VerifyToken(data *pb.TokenRequest) (res *pb.TokenReply, err error)
}
