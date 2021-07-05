package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) CreateMoney(ctx context.Context, input *userin.CreateMoneyInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "validate error", err
	}
	user := input.CreateMoneyInputToUserDomain()

	err = impl.repoMoney.Create(ctx, user, user.Value)
	if err != nil {
		return "error create in money", err
	}
	return user.ID, nil
}
