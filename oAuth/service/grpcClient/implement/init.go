package implement

import (
	"github.com/gnnchya/PosCoffee/oAuth/service/grpcClient"
	"github.com/gnnchya/PosCoffee/oAuth/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/oAuth/service/util"
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
