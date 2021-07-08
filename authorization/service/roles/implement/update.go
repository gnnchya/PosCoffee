package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *rolesin.UpdateInput) (err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return util.ValidationUpdateErr(err)
	}

	filters := makeRoleIDFilters(input.ID)
	roles := &domain.Roles{}
	err = impl.repo.Read(ctx, filters, roles)
	if err != nil {
		return util.RepoReadErr(err)
	}

	update := input.ToDomain()
	update.ID = ""
	err = impl.repo.Update(ctx, filters, update)
	if err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}