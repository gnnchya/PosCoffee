package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
)

func (impl *implementation) Create(ctx context.Context, input *userin.Input) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	initID := goxid.New()
	input.ID = initID.Gen()
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)
	//TODO grpc request for menu here
	err = impl.repo.Create(ctx, user)

	if err != nil {
		return "", err
	}

	return input.ID, nil
}
