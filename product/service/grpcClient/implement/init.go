package implement

import (
	"github.com/gnnchya/PosCoffee/product/service/grpcClient"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
	"github.com/gnnchya/PosCoffee/product/service/util"
	"google.golang.org/grpc"
)

type implementation struct {
	conn      *grpc.ClientConn
	client 	  protobuf.SendIngredientsClient
	connReport *grpc.ClientConn
	clientReport pb.SendReportToStockClient
}

func New(grpcRepo util.RepositoryGRPC, grpcReportRepo util.RepositoryReportGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}

	connReport, err := grpcReportRepo.NewClient()
	if err != nil {
		return nil
	}

	impl := &implementation{
		conn:      conn,
		client: protobuf.NewSendIngredientsClient(conn),
		connReport: connReport,
		clientReport: pb.NewSendReportToStockClient(connReport),
	}
	return impl
}

