package implement

import (
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/cart/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	conn      *grpc.ClientConn
	client protobuf.SendCartClient
	connMiddleware *grpc.ClientConn
	clientMiddleware protobuf.MiddlewareClient
}

func New(grpcRepo util.RepositoryGRPC, grpcRepoMiddleware util.RepositoryGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}
	connMiddleware, err := grpcRepoMiddleware.NewClient()
	if err != nil{
		return nil
	}
	impl := &implementation{
		conn:      conn,
		client: protobuf.NewSendCartClient(conn),
		connMiddleware: connMiddleware,
		clientMiddleware: protobuf.NewMiddlewareClient(connMiddleware),
	}
	return impl

}
