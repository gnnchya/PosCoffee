package implement

import (
	"github.com/gnnchya/PosCoffee/menu/service/grpcClient"
	pb "github.com/gnnchya/PosCoffee/menu/service/grpcClient/protobuf/cart"
	"github.com/gnnchya/PosCoffee/menu/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	conn   *grpc.ClientConn
	client pb.GetMenuClient
}

func New(grpcRepo util.RepositoryGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}
	impl := &implementation{
		conn: conn,
		client: pb.NewGetMenuClient(conn),
	}
	return impl

}
