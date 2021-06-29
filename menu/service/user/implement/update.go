package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
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
	err = impl.elasRepo.Update(ctx, user)
	if err != nil {
		return "", err
	}

	_ = impl.redisRepo.Del(ctx, user.ID)
	return user.ID, nil
}
