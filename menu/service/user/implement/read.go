package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/domain"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ReadInput) (a domain.CreateStruct, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return a, err
	}
	user := userin.ViewInputToUserDomain(input)
	fmt.Println("user input view: ", user)
	fmt.Println("user.ID: ", user.ID)

	if err := impl.redisRepo.Get(ctx, user.ID, a); err == nil{
		return a , nil
	}

	a, err = impl.elasRepo.Read(user.ID, ctx)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	// TODO message broker to redis
	_ = impl.redisRepo.Set(ctx, user.ID, a)
	return a, nil
}
