package implement

import (
	grpcClient "github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	"github.com/gnnchya/PosCoffee/cart/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	conn      *grpc.ClientConn
	conn2 *grpc.ClientConn
}

func New(grpcRepo util.RepositoryGRPC, grpcOAuthRepo util.RepositoryGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}

	conn2, err := grpcOAuthRepo.NewClient()
	if err != nil {
		return nil
	}

	impl := &implementation{
		conn:      conn,
		conn2: conn2,
	}
	return impl

}
