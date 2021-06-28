package implement

import (
	"github.com/gnnchya/PosCoffee/product/service/grpcClient"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	conn      *grpc.ClientConn
	client 	  protobuf.SendIngredientsClient
}

func New(grpcRepo util.RepositoryGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}
	impl := &implementation{
		conn:      conn,
		client: protobuf.NewSendIngredientsClient(conn),
	}
	return impl
}

