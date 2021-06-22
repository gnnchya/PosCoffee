package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Search(ctx context.Context, input *userin.Search) (result string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "", err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.repo.Search(ctx, user)
	fmt.Println("output search:", user)
	if err != nil {
		return a, err
	}

	return a, nil
}
