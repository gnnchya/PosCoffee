package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	goxid "github.com/touchtechnologies-product/xid"
)

func (impl *implementation) Create(ctx context.Context, input *domain.CreateStruct) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	initID := goxid.New()
	input.ID = initID.Gen()
	//user := userin.CreateInputToUserDomain(input)
	//user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", input)
	//TODO grpc receive menu here
	err = impl.repo.Create(ctx, input)

	if err != nil {
		return "", err
	}

	return input.ID, nil
}
