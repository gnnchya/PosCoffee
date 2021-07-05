package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)

	err = impl.repo.Create(ctx, user, user.ID)
	fmt.Println("err in create", err)
	if err != nil {
		return "error create in stock", err
	}
	return user.ID, nil
}