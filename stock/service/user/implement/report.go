package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

func (impl *implementation) Report(ctx context.Context) ([]domain.Report, error) {
	return 	impl.repo.Report(ctx)

}

func (impl *implementation) ReportStock(ctx context.Context, field string, order string) ([]domain.CreateStruct, error) {
	return impl.repo.ReadAllSorted(ctx, field, order)
}