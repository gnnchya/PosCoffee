package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/createFile"
	"github.com/gnnchya/PosCoffee/product/service/reportSale"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"strconv"
	"time"
)

func(impl *implementation)ReportSale(ctx context.Context, input *userin.ReportRange){
	total, _ := impl.repo.ReadMenuTotalSale(ctx,input.From,input.Until)
	fromYear, fromMonth , fromDate:= time.Unix(input.From,0).Date()
	untilYear, untilMonth , untilDate:= time.Unix(input.From,0).Date()
	filename := "./report/reportSale-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)
	switch input.Format{
	case "excel":
		createFile.CreateExcel(filename+".xlsx",reportSale.ReportSale(total))
	case "csv":
		createFile.CreateCSV(filename+".csv",reportSale.ReportSale(total))
	}
}
