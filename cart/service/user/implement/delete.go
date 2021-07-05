package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	user := userin.DeleteInputToUserDomain(input)
	err = impl.repo.Delete(ctx, user.ID)
	if err != nil {
		return "", err
	}
	return user.ID, err
}


