package grpcClient

import pb "github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf/authen"

type Service interface {
	VerifyToken(data *pb.TokenRequest) (res *pb.TokenReply, err error)
}

