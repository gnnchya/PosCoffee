package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {

	err = impl.validator.Validate(input)
	if err != nil {
		return "validate error", err
	}
	user := input.UpdateInputToUserDomain()
	err = impl.repo.Update(ctx, user, user.ID)
	if err != nil {
		return "", err
	}

	return ID, nil
}
