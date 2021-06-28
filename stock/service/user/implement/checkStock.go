package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

func (impl *implementation) CheckStock(ctx context.Context, input []string) (state bool, expenses []domain.CalculateCost, err string) {
	state, expenses, e := impl.repo.CheckMenuAvailability(ctx, input)
	err = e.Error()
	return state, expenses, err
}