package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/domain"
)

func (impl *implementation) ReadMoneyAll(ctx context.Context)([]domain.CreateMoneyStruct, error) {
	a, err := impl.repoMoney.ReadMoneyAll(ctx)
	if err != nil {
		return a, err
	}
	return a, nil
}
