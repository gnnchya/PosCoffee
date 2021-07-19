package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) ReadMoney(ctx context.Context, input *userin.ReadMoneyInput) (a interface{}, err error) {
	user := input.ReadMoneyInputToUserDomain()

	a, err = impl.repoMoney.Read(ctx, user.ID)

	if err != nil {
		return a, err
	}

	return a, nil
}
