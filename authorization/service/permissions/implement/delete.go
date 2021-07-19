package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *permissionsin.DeleteInput) (err error) {
	permission := &domain.Permissions{}
	filters := makePermissionIDFilters(input.ID)

	err = impl.repo.Read(ctx, filters, permission)
	if err != nil {
		return util.RepoReadErr(err)
	}

	err = impl.repo.SoftDelete(ctx, filters)
	if err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
