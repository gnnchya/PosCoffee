package implement

import (
	"context"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
	"time"
)

func (impl *implementation) SendReportToStock(input *pb.ReportRequest) (*pb.ReportReply, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	r, err := impl.clientReport.SendReportToStock(ctx, input)
	if err != nil {
		return r, err
	}
	return r, nil
}