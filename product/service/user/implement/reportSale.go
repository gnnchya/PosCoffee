package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/reportSale"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func(impl *implementation)ReportSale(ctx context.Context, input *userin.ReportRange) ([][]string, error){
	total, err := impl.repo.ReadMenuTotalSale(ctx,input.From,input.Until)
	//menu := //proud
	request := &protobuf.RequestMenu{Err: "nil"}
	reply, err := impl.client.SendMenu(request)
	if err != nil{
		return [][]string{}, err
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
	rangeReport := reportSale.ReportSale(total,menu)
	fmt.Println("range report", rangeReport)
	return rangeReport, err
}
