package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	user := userin.DeleteInputToUserDomain(input)

	err = impl.repo.Delete(ctx, ID)
	if err != nil {
		return "err delete in stock", err
	}
	return user.ID, err
}


