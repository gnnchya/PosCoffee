package implement

import (
	grpcService "github.com/gnnchya/PosCoffee/product/service/grpc"
	"github.com/gnnchya/PosCoffee/product/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/user"
	"github.com/gnnchya/PosCoffee/product/service/util"
	"google.golang.org/grpc"
	"log"
)

type implementation struct {
	userService user.Service
}


const(
	NETWORK = "tcp"
)

func New(grpcRepo util.RepositoryGRPC, userService user.Service) (service grpcService.Service){
	impl := implementation{
		userService: userService,
	}

	//lis, err := net.Listen(NETWORK, conf.GRPCHost)
	grpcServer := grpc.NewServer()
	lis, err := grpcRepo.NetListener()
	protobuf.RegisterSendCartServer(grpcServer, &impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return impl
}
