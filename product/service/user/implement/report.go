package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/createFile"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
	"github.com/gnnchya/PosCoffee/product/service/report"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"strconv"
	"time"
)

func(impl *implementation)Report(ctx context.Context, input *userin.ReportRange){
	transaction,_ := impl.repo.ReadByTimeRange(ctx, input.From,input.Until)
	out := &pb.ReportRequest{Request: "stock"}
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
	fromYear, fromMonth , fromDate:= time.Unix(input.From,0).Date()
	untilYear, untilMonth , untilDate:= time.Unix(input.From,0).Date()
	filename := "./report/report-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)
	switch input.Format{
	case "excel":
		createFile.CreateExcel(filename+".xls",report.Report(transaction, stock))
	case "csv":
		createFile.CreateCSV(filename+".csv",report.Report(transaction, stock))
	}
}
