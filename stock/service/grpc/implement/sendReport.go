package implement

import (
	"context"
	"fmt"
	pb "github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf/report"
)

func(impl implementation) SendReportToStock(ctx context.Context, input *pb.ReportRequest) (*pb.ReportReply, error){
	fmt.Println("from product ask for report", input)
	report, err := impl.userService.Report(ctx)
	fmt.Println("report", report)
	if err != nil {
		return nil, err
	}
	var result []*pb.ReportReplyStruct
	for _, v := range report{
		temp := &pb.ReportReplyStruct{
			ID:          v.ID,
			ItemName:    v.ItemName,
			Category:    v.Category,
			Amount:      v.Amount,
			Unit:        v.Unit,
			CostPerUnit: v.CostPerUnit,
			EXPDate:     v.EXPDate,
			ImportDate:  v.ImportDate,
			Supplier:    v.Supplier,
			TotalCost:   v.TotalCost,
			TotalAmount: v.TotalAmount,
			Status:      v.Status,
		}
		result = append(result, temp)
	}

	fmt.Println("result", result)
	out := &pb.ReportReply{Report: result}
	return out, nil
}