package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/uniplaces/carbon"
)

func (impl *implementation) SoftDelete(ctx context.Context, input *userin.DeleteInput) (err error) {
	user := &domain.UserStruct{}
	filters := impl.filter.MakeIdFilters(input.ID)

	err = impl.repo.Read(ctx, filters, user)
	if err != nil {
		return err
	}

	update := &domain.UsersSoftDeleteStruct{
		DeletedAt: carbon.Now().Unix(),
	}

	err = impl.repo.Update(ctx, filters, update)
	if err != nil {
		return err
	}

	return nil
}

