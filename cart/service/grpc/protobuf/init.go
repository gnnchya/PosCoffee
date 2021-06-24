package protobuf

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/config"
	pb "github.com/gnnchya/PosCoffee/cart/service/grpc/protobuf/cart"
	"github.com/gnnchya/PosCoffee/cart/service/user"
	"google.golang.org/grpc"
	"net"
)

type GetMenuServerImpl struct {
	userService user.Service
	pb.UnimplementedGetMenuServer
}


const(
	NETWORK = "tcp"
)

func NewServer(conf *config.Config, userService user.Service) {
	server := GetMenuServerImpl{
		userService: userService,
	}

	lis, err := net.Listen(NETWORK, conf.GRPCHost)

	grpcServer := grpc.NewServer()
	pb.RegisterGetMenuServer(grpcServer, &server)

	// start grpcServer
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func (impl *GetMenuServerImpl) Menu(ctx context.Context, data *pb.Request) (*pb.Reply, error) {
	resp := &pb.Reply{
		Id:  data.Id,
		Arr: "hello menu",
	}
	return resp, nil
}
