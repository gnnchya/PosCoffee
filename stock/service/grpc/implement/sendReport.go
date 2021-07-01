package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	pb "github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf/report"
)

func(impl implementation) SendReportToStock(ctx context.Context, input *pb.ReportRequest) (*pb.ReportReply, error){
	fmt.Println("from product ask for report", input)
	var err error
	var report []domain.CreateStruct
	if input.Request == "stock"{
		report, err = impl.userService.Report(ctx)
	}else if input.Request == "reportStock"{
		report, err = impl.userService.ReportStock(ctx, input.Field, input.Order)
	}
	//report, err := impl.userService.Report(ctx)
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