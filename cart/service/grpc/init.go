package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/config"
	"github.com/gnnchya/PosCoffee/cart/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/cart/service/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type implementation struct {
	userService user.Service
}


const(
	NETWORK = "tcp"
)

func NewServer(conf *config.Config, userService user.Service){
	impl := implementation{
		userService: userService,
	}

	lis, err := net.Listen(NETWORK, conf.GRPCHost)
	grpcServer := grpc.NewServer()
	//lis, err := grpcRepo.NetListener()
	protobuf.RegisterSendCartServer(grpcServer, &impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (impl *implementation)ReceiveChange(ctx context.Context, data *protobuf.Request) (*protobuf.Reply, error){
	return nil, nil
}