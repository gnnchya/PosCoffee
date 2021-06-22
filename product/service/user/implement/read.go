package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ReadInput) (a interface{}, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return nil, err
	}
	user := userin.ReadInputToUserDomain(input)

	a, err = impl.repo.Read(ctx, user.ID)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return a, nil
}
