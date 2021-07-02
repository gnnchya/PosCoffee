package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/createFile"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
	"github.com/gnnchya/PosCoffee/product/service/reportStock"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"strconv"
	"time"
)

func(impl *implementation)ReportStock(ctx context.Context, input *userin.ReportFilter){
	//user := input.ReportStockInputToUserDomain()
	out := &pb.ReportRequest{
		Request: "reportStock",
		Field:   input.Field,
		Order:   input.Order,
	}
	reply, err := impl.client.SendReportToStock(out)
	fmt.Println("reply", reply)
	if err != nil{
		return
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
	fromYear, fromMonth , fromDate:= time.Now().Date()
	filename := "./report/reportStock-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)
	switch input.Format{
	case "excel":
		createFile.CreateExcel(filename+".xls",reportStock.ReportStock(stock))
	case "csv":
		createFile.CreateCSV(filename+".csv",reportStock.ReportStock(stock))
	}
}
