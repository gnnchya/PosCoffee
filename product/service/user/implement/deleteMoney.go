package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) DeleteMoney(ctx context.Context, input *userin.DeleteMoneyInput) (ID string, err error) {

	user := input.DeleteMoneyInputToUserDomain()
	fmt.Println("user input delete:", user)

	err = impl.repoMoney.Delete(ctx, user.ID)

	if err != nil {
		return "", err
	}
	return ID, err
}
