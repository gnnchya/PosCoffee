package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
	"github.com/gnnchya/PosCoffee/product/service/report"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func(impl *implementation)Report(ctx context.Context, input *userin.ReportRange) ([][]string, error){
	transaction,_ := impl.repo.ReadByTimeRange(ctx, input.From,input.Until)
	//stock := //proud
	_,_ = impl.repo.ReadMenuTotalSale(ctx,input.From,input.Until)
	out := &pb.ReportRequest{Request: "stock"}
	reply, err := impl.client.SendReportToStock(out)
	fmt.Println("reply", reply)
	if err != nil{
		return nil, err
	}
	var stock []domain.CreateStockStruct
	for _, v := range reply.Report{
		temp := domain.CreateStockStruct{
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
		stock = append(stock, temp)
	}
	rangeReport := report.Report(transaction, stock)
	return rangeReport, nil
}
