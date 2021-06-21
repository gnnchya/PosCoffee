package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	// "github.com/touchtechnologies-product/go-blueprint-clean	-architecture/service/company/companyin"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) ReadCategoryAll(ctx context.Context, input *userin.ReadCategoryAllInput)([]domain.CreateStruct, error) {
	user := userin.ReadCategoryInputToUserDomain(input)
	//TODO change to use repo from mongodb
	a, err := impl.elasRepo.ViewAll(user.Page, user.PerPage,ctx)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (impl *implementation) ReadNameAll(ctx context.Context, input *userin.ReadNameAllInput)([]domain.CreateStruct, error) {
	user := userin.ReadNameInputToUserDomain(input)
	//TODO change to use repo from mongodb
	a, err := impl.elasRepo.ViewAll(user.Page, user.PerPage,ctx)
	if err != nil {
		return a, err
	}

	return a, nil
}
