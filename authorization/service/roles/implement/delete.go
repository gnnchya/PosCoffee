package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *rolesin.DeleteInput) (err error) {
	role := &domain.Roles{}
	filters := makeRoleIDFilters(input.ID)

	err = impl.repo.Read(ctx, filters, role)
	if err != nil {
		return util.RepoReadErr(err)
	}

	err = impl.repo.SoftDelete(ctx, filters)
	if err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
