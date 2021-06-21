package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) SearchMenu(ctx context.Context, input *userin.Search) ([]domain.InsertStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.elasRepo.SearchMenu(user.Value, ctx)
	fmt.Println("output search:", user)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (impl *implementation) SearchIngredient(ctx context.Context, input *userin.Search) ([]domain.InsertStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.elasRepo.SearchIngredient(user.Value, ctx)
	fmt.Println("output search:", user)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (impl *implementation) SearchCategory(ctx context.Context, input *userin.Search) ([]domain.InsertStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.SearchInputToUserDomain(input)
	fmt.Println("user input search:", user)
	a, err := impl.elasRepo.SearchCategory(user.Value, ctx)
	fmt.Println("output search:", user)
	if err != nil {
		return a, err
	}

	return a, nil
}
