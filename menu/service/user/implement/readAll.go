package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/domain"

	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) ReadAll(ctx context.Context, input *userin.ViewAllInput)([]domain.CreateStruct, error) {
	// err = impl.validator.Validate(input)
	// if err != nil {
	// 	return "", err
	// }

	user := userin.ViewAllInputToUserDomain(input)

	a, err := impl.elasRepo.ReadAll(user.Page, user.PerPage,ctx)
	if err != nil {
		return a, err
	}

	return a, nil
}
