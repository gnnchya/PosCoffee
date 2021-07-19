package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/domain"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
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
	var a domain.CreateStruct
	if err := impl.redisRepo.Get(ctx, user.ID, &a); err != nil{
		fmt.Println("delete successfully Redis: or error" , err)
	}
	return user.ID, nil
}
