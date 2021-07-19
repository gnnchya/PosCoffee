package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient"
	"github.com/gnnchya/PosCoffee/authen/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	conn      *grpc.ClientConn
	client protobuf.AuthorizeClient
}

func New(grpcRepo util.RepositoryGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}
	impl := &implementation{
		conn:      conn,
		client: protobuf.NewAuthorizeClient(conn),
	}
	return impl

}
