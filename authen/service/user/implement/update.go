package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {

	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}

	user := input.UpdateInputToUserDomain()
	fmt.Println("user in service update", user)
	err = impl.userRepo.Update(ctx, user, user.ID)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
