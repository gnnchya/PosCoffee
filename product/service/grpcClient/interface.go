package grpcClient

import (
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
)

type Service interface {
	SendIngredients(input *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error)
	SendReportToStock(input *pb.ReportRequest) (*pb.ReportReply, error)
	ReadStock(input *protobuf.RequestRead) (*protobuf.ReplyRead, error)
	SendMenu(input *protobuf.RequestMenu) (*protobuf.ReplyMenu, error)
	ReadNameStock(input *protobuf.RequestName)(*protobuf.ReplyArrRead, error)
	ReadCategoryStock(input *protobuf.RequestCategory)(*protobuf.ReplyArrRead, error)
	Middleware(input *protobuf.RequestMiddleware) (*protobuf.ReplyMiddleware, error)

}

