package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ReadInput) (a domain.Users, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return a, err
	}
	user := userin.ViewInputToUserDomain(input)
	fmt.Println("user input view: ", user)
	fmt.Println("user.ID: ", user.ID)

	a, err = impl.userRepo.Read(ctx, user.ID)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return a, nil
}
