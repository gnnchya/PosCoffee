package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"

	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ViewInput) (domain.CreateStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		return domain.CreateStruct{}, err
	}
	user := userin.ViewInputToUserDomain(input)
	fmt.Println("userid", user.CustomerID)
	a, err := impl.repo.Read(ctx,user.CustomerID)
	fmt.Println("a", a)
	fmt.Println("err", err)
	if err != nil {
		return a, err
	}

	return a, nil
}
