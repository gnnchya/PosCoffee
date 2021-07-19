package grpc

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"
	pb "github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf/report"
)

type Service interface{
	SendIngredients(context.Context, *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error)
	SendReportToStock(ctx context.Context, input *pb.ReportRequest) (*pb.ReportReply, error)
	ReadStock(ctx context.Context, input *protobuf.RequestRead) (*protobuf.ReplyRead, error)
	ReadCategoryStock(context.Context, *protobuf.RequestCategory) (*protobuf.ReplyArrRead, error)
	ReadNameStock(context.Context, *protobuf.RequestName) (*protobuf.ReplyArrRead, error)
}
