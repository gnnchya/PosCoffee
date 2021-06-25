package implement

import (
	pb "github.com/gnnchya/PosCoffee/cart/service/grpc/protobuf/cart"
	grpcService "github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	"github.com/gnnchya/PosCoffee/cart/service/user"
	"github.com/gnnchya/PosCoffee/cart/service/util"
	"google.golang.org/grpc"
	"log"
)

type implementation struct {
	userService user.Service
	pb.UnimplementedGetMenuServer
}


const(
	NETWORK = "tcp"
)

func New(grpcRepo util.RepositoryGRPC, userService user.Service) (service grpcService.Service){
	impl := &implementation{
		userService: userService,
	}

	//lis, err := net.Listen(NETWORK, conf.GRPCHost)

	grpcServer := grpc.NewServer()
	lis, err := grpcRepo.NetListener()
	pb.RegisterGetMenuServer(grpcServer, impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return impl
}