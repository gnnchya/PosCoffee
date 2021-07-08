package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *permissionsin.UpdateInput) (err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return util.ValidationUpdateErr(err)
	}

	filters := makePermissionIDFilters(input.ID)
	permissions := &domain.Permissions{}
	err = impl.repo.Read(ctx, filters, permissions)
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
