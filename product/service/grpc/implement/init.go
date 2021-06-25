package implement

import (
	"github.com/gnnchya/PosCoffee/product/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/util"
	grpcService "github.com/gnnchya/PosCoffee/product/service/grpc"
	userService "github.com/gnnchya/PosCoffee/product/service/user"
	"log"
	"google.golang.org/grpc"
)

type implementation struct {
	userService userService.Service
}
func New(grpcRepo util.RepositoryGRPC, userService userService.Service)(service grpcService.Service){
	impl := &implementation{
		userService: userService,
	}
	grpcServer := grpc.NewServer()
	lis, err := grpcRepo.NetListener()
	protobuf.RegisterSendCartServer(grpcServer, impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return impl
}

