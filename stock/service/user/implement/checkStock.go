package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
func (impl *implementation) CheckStock(ctx context.Context, input []string) (state bool, expenses []domain.CalculateCost, err string) {
	state, expenses, e := impl.repo.CheckMenuAvailability(ctx, input)

	errString := e.Error()
	return state, expenses, errString
}