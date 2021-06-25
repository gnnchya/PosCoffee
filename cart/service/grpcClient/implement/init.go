package implement

import (
	"github.com/gnnchya/PosCoffee/cart/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	"github.com/gnnchya/PosCoffee/cart/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	conn      *grpc.ClientConn
	client protobuf.SendCartClient
}

func New(grpcRepo util.RepositoryGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}
	impl := &implementation{
		conn:      conn,
		client: protobuf.NewSendCartClient(conn),
	}
	return impl

}
