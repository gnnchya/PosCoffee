package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	pb "github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf/report"
	"github.com/gnnchya/PosCoffee/product/service/report"
	"github.com/gnnchya/PosCoffee/product/service/reportSale"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func(impl *implementation)ReportSale(ctx context.Context, input *userin.ReportRange) ([][]string, error){
	total, err := impl.repo.ReadMenuTotalSale(ctx,input.From,input.Until)
	//menu := //proud
	rangeReport := reportSale.ReportSale(total,menu)
	return rangeReport, err
}
