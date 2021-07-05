package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/domain"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	user := userin.DeleteInputToUserDomain(input)

	err = impl.elasRepo.Delete(ctx, user.ID)
	if err != nil {
		return "", err
	}

	_ = impl.redisRepo.Del(ctx, user.ID)
	var a domain.CreateStruct
	err = impl.redisRepo.Get(ctx, user.ID, &a)
	if err != nil {
		return "", err
	}
	return user.ID, err
}


