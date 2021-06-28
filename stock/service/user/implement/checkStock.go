package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

func (impl *implementation) CheckStock(ctx context.Context, input []string) (state bool, expenses []domain.CalculateCost, err string) {
	fmt.Println("input in checkstock service", input)
	state, expenses, e := impl.repo.CheckMenuAvailability(ctx, input)
	err = e.Error()
	return state, expenses, err
}