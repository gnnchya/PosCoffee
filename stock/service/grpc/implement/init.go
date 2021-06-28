package implement

import (
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/stock/service/util"
	"github.com/gnnchya/PosCoffee/stock/service/user"
	grpcService "github.com/gnnchya/PosCoffee/stock/service/grpc"
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
	protobuf.RegisterSendIngredientsServer(grpcServer, &impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return impl
}
