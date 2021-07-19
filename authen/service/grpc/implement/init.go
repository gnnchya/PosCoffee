package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	grpcService "github.com/gnnchya/PosCoffee/authen/service/grpc"
	"github.com/gnnchya/PosCoffee/authen/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/user"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"google.golang.org/grpc"
	"log"
)

type implementation struct {
	userService user.Service
	authService authentication.Service
	repo util.Repository
	filter util.Filters
}

func New(grpcRepo util.RepositoryGRPC, userService user.Service, authService authentication.Service, repo util.Repository, filter util.Filters) (service grpcService.Service){
	impl := implementation{
		userService: userService,
		authService: authService,
		repo:	repo,
		filter: filter,
	}

	grpcServer := grpc.NewServer()
	lis, err := grpcRepo.NetListener()
	protobuf.RegisterMiddlewareServer(grpcServer, &impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return impl
}
