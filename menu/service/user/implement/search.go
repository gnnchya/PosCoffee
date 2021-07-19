package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/domain"

	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) SearchMenu(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.elasRepo.SearchMenu(user.Value, ctx)
	fmt.Println("output search:", a)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (impl *implementation) SearchIngredient(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.elasRepo.SearchIngredient(user.Value, ctx)
	fmt.Println("output search:", a)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (impl *implementation) SearchCategory(ctx context.Context, input *userin.Search) ([]domain.CreateStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.elasRepo.SearchCategory(user.Value, ctx)
	fmt.Println("output search:", a)
	if err != nil {
		return a, err
	}

	return a, nil
}
