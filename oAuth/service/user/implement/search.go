package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"

	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (impl *implementation) Search(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	a, err := impl.repo.Search(ctx, user)
	if err != nil {
		return a, err
	}

	return a, nil
}
