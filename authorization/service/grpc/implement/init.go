package implement

import (
	"github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf"
	"google.golang.org/grpc"

	grpcService "github.com/gnnchya/PosCoffee/authorize/service/grpc"
	pb "github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf/authen"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

type implementation struct {
	conn   *grpc.ClientConn
	client pb.AuthenticationClient
	clientPermission protobuf.AuthorizeClient
}

func New(grpcRepo util.RepositoryGRPC) (service grpcService.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}
	impl := &implementation{
		conn:   conn,
		client: pb.NewAuthenticationClient(conn),
		clientPermission: protobuf.NewAuthorizeClient(conn),
	}
	return impl

}