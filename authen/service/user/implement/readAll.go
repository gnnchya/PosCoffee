package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/domain"

	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) ReadAll(ctx context.Context, input *userin.ViewAllInput)([]domain.Users, error) {

	user := userin.ViewAllInputToUserDomain(input)
	fmt.Println("user in read all:", user)
	a, err := impl.userRepo.ReadAll(ctx, user.Page, user.PerPage)
	if err != nil {
		return a, err
	}

	return a, nil
}