package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"

	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (impl *implementation) ReadAll(ctx context.Context, input *userin.ViewAllInput)(a []domain.CreateStruct, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil , err
	}

	user := userin.ViewAllInputToUserDomain(input)

	a, err = impl.repo.ReadAll(ctx, user.PerPage,user.Page)
	if err != nil {
		return a, err
	}

	return a, nil
}
