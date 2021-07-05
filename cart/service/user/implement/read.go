package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/domain"

	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ViewInput) (domain.CreateStruct, error) {
	err := impl.validator.Validate(input)
	if err != nil {
		return domain.CreateStruct{}, err
	}
	user := userin.ViewInputToUserDomain(input)
	a, err := impl.repo.Read(ctx,user.ID)
	if err != nil {
		return a, err
	}

	return a, nil
}
