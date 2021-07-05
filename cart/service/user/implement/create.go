package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
)

func (impl *implementation) Create(ctx context.Context, input *userin.Input) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "validate error", err
	}
	initID := goxid.New()
	input.ID = initID.Gen()
	initID = goxid.New()
	input.CustomerID = initID.Gen()
	user := input.CreateInputToUserDomain()

	err = impl.repo.Create(ctx, user)

	if err != nil {
		return "", err
	}

	return input.ID, nil
}
