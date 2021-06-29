package grpcClient

import (
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
)

type Service interface {
	SendIngredients(input *protobuf.RequestToStock) (*protobuf.ReplyFromStock, error)
	SendReportToStock(input *pb.ReportRequest) (*pb.ReportReply, error)
}

