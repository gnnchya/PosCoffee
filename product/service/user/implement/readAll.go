package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) ReadAll(ctx context.Context, input *userin.ReadAllInput)([]interface{}, error) {
	// err = impl.validator.Validate(input)
	// if err != nil {
	// 	return "", err
	// }

	user := userin.ReadAllInputToUserDomain(input)

	a, err := impl.repo.ReadOrderAll(ctx, user)
	if err != nil {
		return a, err
	}

	return a, nil
}
