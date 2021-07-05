package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) UpdateMoney(ctx context.Context, input *userin.UpdateMoneyInput) (ID string, err error) {

	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}

	user := input.UpdateStockInputToUserDomain()

	err = impl.repoMoney.Update(ctx, user, user.ID)
	if err != nil {
		return "", err
	}

	return ID, nil
}
