package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {

	user := userin.DeleteInputToUserDomain(input)

	err = impl.repo.Delete(ctx, user.ID)

	if err != nil {
		return "", err
	}
	return ID, err
}

