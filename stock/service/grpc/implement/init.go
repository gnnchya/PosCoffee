package implement

import (
	grpcService "github.com/gnnchya/PosCoffee/stock/service/grpc"
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"
	pb "github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf/report"
	"github.com/gnnchya/PosCoffee/stock/service/user"
	"github.com/gnnchya/PosCoffee/stock/service/util"
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
	protobuf.RegisterReadStockServer(grpcServer, &impl)
	pb.RegisterSendReportToStockServer(grpcServer, &impl)
	protobuf.RegisterReadCategoryStockServer(grpcServer, &impl)
	protobuf.RegisterReadNameStockServer(grpcServer, &impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return impl
}
