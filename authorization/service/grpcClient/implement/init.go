package implement

import (
	"google.golang.org/grpc"

	pb "github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf/authen"
	grpcService "github.com/gnnchya/PosCoffee/authorize/service/grpcClient"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

type implementation struct {
	conn   *grpc.ClientConn
	client pb.AuthenticationClient
}

func New(grpcRepo util.RepositoryGRPC) (service grpcService.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}
	impl := &implementation{
		conn:   conn,
		client: pb.NewAuthenticationClient(conn),
	}
	return impl

}