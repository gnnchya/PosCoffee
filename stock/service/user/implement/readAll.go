package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) ReadCategoryAll(ctx context.Context, input *userin.ReadCategoryAllInput)([]domain.CreateStruct, error) {
	user := userin.ReadCategoryInputToUserDomain(input)
	fmt.Println("user input in service read category:", user)
	a, err := impl.repo.ReadCategoryAll(ctx, user)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (impl *implementation) ReadNameAll(ctx context.Context, input *userin.ReadNameAllInput)([]domain.CreateStruct, error) {
	user := userin.ReadNameInputToUserDomain(input)
	a, err := impl.repo.ReadNameAll(ctx, user)
	if err != nil {
		return a, err
	}
	return a, nil
}
