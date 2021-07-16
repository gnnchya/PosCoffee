package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	user := userin.DeleteInputToUserDomain(input)
	_ = impl.redisRepo.Del(ctx, user.ID)

	err = impl.elasRepo.Delete(ctx, user.ID)
	if err != nil {
		return "", err
	}
	return user.ID, err
}


