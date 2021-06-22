package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ViewInput) ([]domain.InsertStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.ViewInputToUserDomain(input)
	fmt.Println("user input view: ", user)
	fmt.Println("user.ID: ", user.ID)
	a, err := impl.elasRepo.Read(user.ID, ctx)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return a, nil
}