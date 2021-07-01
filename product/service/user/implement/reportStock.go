package implement

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
	"github.com/gnnchya/PosCoffee/product/service/reportStock"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func(impl *implementation)ReportStock(input *userin.ReportFilter) ([][]string, error){
	user := input.ReportStockInputToUserDomain()
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
	rangeReport := reportStock.ReportStock(stock)
	return rangeReport, nil
}
