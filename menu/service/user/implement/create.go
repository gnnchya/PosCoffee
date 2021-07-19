package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)

	err = impl.elasRepo.Create(ctx, input)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
