package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/report"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"os"
)

func(impl *implementation)Report(ctx context.Context, input *userin.ReportRange) *os.File{
	transaction,_ := impl.repo.ReadByTimeRange(ctx, input.From,input.Until)
	//stock := //proud
	var stock []domain.CreateStockStruct
	report := report.Report(transaction, stock)
	return report
}
