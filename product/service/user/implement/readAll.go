package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) ReadAll(ctx context.Context, input *userin.ReadAllInput)([]domain.CreateOrderStruct, error) {
	user := userin.ReadAllInputToUserDomain(input)

	a, err := impl.repo.ReadOrderAll(ctx, user)
	if err != nil {
		return a, err
	}

	return a, nil
}
