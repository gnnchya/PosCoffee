package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/createFile"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/reportSale"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"strconv"
	"time"
)

func(impl *implementation)ReportSale(ctx context.Context, input *userin.ReportRange){
	total, err := impl.repo.ReadMenuTotalSale(ctx,input.From,input.Until)
	//menu := //proud
	request := &protobuf.RequestMenu{Err: "nil"}
	reply, err := impl.client.SendMenu(request)
	if err != nil{
		return
	}
	fmt.Println("menu from menu:", reply)
	var menu []domain.OldMenu
	for _, k := range reply.Menu{
		var ingre []domain.Ingredient

		for _, j := range k.Ingredient{
			in := domain.Ingredient{
				IngredientName: j.IngredientName,
				Amount:         j.Amount,
			}
			ingre = append(ingre, in)
		}
		m := domain.OldMenu{
			ID:         k.ID,
			Category:   k.Category,
			Name:       k.Name,
			Ingredient: ingre,
			Price:      k.Price,
			Available:  k.Available,
		}
		menu = append(menu, m)
	}
	fromYear, fromMonth , fromDate:= time.Unix(input.From,0).Date()
	untilYear, untilMonth , untilDate:= time.Unix(input.From,0).Date()
	filename := "./report/reportSale-"+strconv.Itoa(fromDate)+"."+strconv.Itoa(int(fromMonth))+"."+strconv.Itoa(fromYear)+"-"+strconv.Itoa(untilDate)+"."+strconv.Itoa(int(untilMonth))+"."+strconv.Itoa(untilYear)
	switch input.Format{
	case "excel":
		createFile.CreateExcel(filename+".xlsx",reportSale.ReportSale(total, menu))
	case "csv":
		createFile.CreateCSV(filename+".csv",reportSale.ReportSale(total, menu))
	}
}
