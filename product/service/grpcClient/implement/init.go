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
	clientRead protobuf.ReadStockClient
	clientReadName protobuf.ReadNameStockClient
	clientReadCategory protobuf.ReadCategoryStockClient
	connReport *grpc.ClientConn
	clientReport pb.SendReportToStockClient
	connMenu *grpc.ClientConn
	clientMenu protobuf.SendMenuClient
}

func New(grpcRepo util.RepositoryGRPC, grpcReportRepo util.RepositoryReportGRPC, grpcMenu util.RepositoryGRPC) (service grpcClient.Service) {
	conn, err := grpcRepo.NewClient()
	if err != nil {
		return nil
	}

	connReport, err := grpcReportRepo.NewClient()
	if err != nil {
		return nil
	}

	connMenu, err := grpcMenu.NewClient()
	if err != nil {
		return nil
	}


	impl := &implementation{
		conn:      conn,
		client: protobuf.NewSendIngredientsClient(conn),
		clientRead: protobuf.NewReadStockClient(conn),
		clientReadName: protobuf.NewReadNameStockClient(conn),
		clientReadCategory: protobuf.NewReadCategoryStockClient(conn),
		connReport: connReport,
		clientReport: pb.NewSendReportToStockClient(connReport),
		connMenu: connMenu,
		clientMenu: protobuf.NewSendMenuClient(connMenu),
	}
	return impl
}

