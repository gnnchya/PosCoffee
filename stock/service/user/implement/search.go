package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"

	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Search(ctx context.Context, input *userin.Search) (a []domain.CreateStruct, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	//a, err := impl.repo.Search(user.Value, ctx)
	fmt.Println("output search:", user)
	if err != nil {
		return a, err
	}

	return a, nil
}
