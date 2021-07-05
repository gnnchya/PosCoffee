package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/domain"

	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) SearchName(ctx context.Context, input *userin.Search) ([]domain.Users, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.userRepo.SearchName(ctx, user)
	fmt.Println("output search:", user)
	if err != nil {
		return a, err
	}

	return a, nil
}

