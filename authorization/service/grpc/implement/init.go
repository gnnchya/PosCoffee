package implement

import (
	"github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf"
	permissionsService "github.com/gnnchya/PosCoffee/authorize/service/permissions"
	rolesService "github.com/gnnchya/PosCoffee/authorize/service/roles"
	"google.golang.org/grpc"
	"log"

	grpcService "github.com/gnnchya/PosCoffee/authorize/service/grpc"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

type implementation struct {
	permissionService permissionsService.Service
	roleService rolesService.Service
}

func New(grpcRepo util.RepositoryGRPC, permission permissionsService.Service, role rolesService.Service) (service grpcService.Service) {
	impl := implementation{
		permissionService: permission,
		roleService:       role,
	}
	grpcServer := grpc.NewServer()
	lis, err := grpcRepo.NetListener()
	protobuf.RegisterAuthorizeServer(grpcServer, &impl)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return impl
}