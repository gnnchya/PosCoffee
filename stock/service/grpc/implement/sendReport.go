package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	pb "github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf/report"
)

func(impl implementation) SendReportToStock(ctx context.Context, input *pb.ReportRequest) ( out *pb.ReportReply, err error){
	fmt.Println("from product ask for report", input)
	if input.Request == "stock"{
		var report []domain.Report
		report, err = impl.userService.Report(ctx)
		fmt.Println("report", report)
		if err != nil {
			return nil, err
		}
		var result []*pb.ReportReplyStruct
		for _, v := range report{
			temp := &pb.ReportReplyStruct{
				ItemName:    v.ID.ItemName,
				Category:    v.ID.Category,
				Amount:      v.Amount,
				Unit:        v.ID.Unit,
				TotalCost:   v.TotalCost,
				TotalAmount: v.TotalAmount,
			}
			result = append(result, temp)
		}

		fmt.Println("result", result)
		out := &pb.ReportReply{Report: result}
		return out, nil
	}else if input.Request == "reportStock"{
		var report []domain.CreateStruct
		report, err = impl.userService.ReportStock(ctx, input.Field, input.Order)
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
	//report, err := impl.userService.Report(ctx)

	return out, nil
}