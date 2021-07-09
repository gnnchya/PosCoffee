package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) Read(ctx context.Context, input *userin.ReadInput) (view *domain.UserStruct, err error) {
	user := &domain.UserStruct{}
	filters := impl.filter.MakeIdFilters(input.ID)

	err = impl.repo.Read(ctx, filters, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
