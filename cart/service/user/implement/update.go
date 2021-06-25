package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"

	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *userin.Input) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	//user := userin.UpdateInputToUserDomain(input)
	user := input.UpdateInputToUserDomain()
	fmt.Println("user update input", user)
	//TODO request menu by grpc
	err = impl.repo.Update(ctx, user, user.ID)
	//time.Sleep(5 * time.Second)
	//_, err = impl.repo.Read(ctx, input.ID)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

