package implement

import (
	"github.com/gnnchya/PosCoffee/menu/service/grpcClient"
	"github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/menu/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	connMiddleware *grpc.ClientConn
	clientMiddleware protobuf.MiddlewareClient
}

func New(grpcRepoMiddleware util.RepositoryGRPC) (service grpcClient.Service) {

	connMiddleware, err := grpcRepoMiddleware.NewClient()
	if err != nil{
		return nil
	}
	impl := &implementation{

		connMiddleware: connMiddleware,
		clientMiddleware: protobuf.NewMiddlewareClient(connMiddleware),
	}
	return impl

}
