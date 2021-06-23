package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	initID := goxid.New()
	input.ID = initID.Gen()
	//user := userin.CreateInputToUserDomain(input)
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)

	err = impl.repo.Create(ctx, user)
	// fmt.Println("output create:", user)

	//if err != nil {
	//	return "", err
	//}

	//time.Sleep(5 * time.Second)
	//_, err = impl.repo.Read(ctx, input.ID)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
