package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) ReadAll(ctx context.Context, input *userin.ReadAllInput)([]interface{}, error) {
	user := userin.ReadAllInputToUserDomain(input)

	a, err := impl.repo.ReadOrderAll(ctx, user)
	if err != nil {
		return a, err
	}

	return a, nil
}
