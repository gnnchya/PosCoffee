package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"

	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ViewInput) (domain.CreateStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return domain.CreateStruct{}, err
	}
	user := userin.ViewInputToUserDomain(input)
	fmt.Println("user input view: ", user)
	fmt.Println("user.ID: ", user.ID)
	a, err := impl.repo.Read(ctx,user.ID)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return a, nil
}
