package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/domain"

	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) ReadAll(ctx context.Context, input *userin.ViewAllInput)([]domain.CreateStruct, error) {

	user := userin.ViewAllInputToUserDomain(input)
	fmt.Println("user in read all:", user)
	a, err := impl.elasRepo.ReadAll(user.Page, user.PerPage,ctx)
	if err != nil {
		return a, err
	}

	return a, nil
}
