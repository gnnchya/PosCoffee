package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ReadInput) (a domain.CreateStruct, err error) {
	user := userin.ReadInputToUserDomain(input)
	a, err = impl.repo.Read(ctx, user.ID)
	fmt.Println("err:", err)
	if err != nil {
		return a, err
	}

	return a, nil
}
