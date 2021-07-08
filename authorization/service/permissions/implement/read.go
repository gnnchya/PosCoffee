package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *permissionsin.ReadInput) (view *out.PermissionsView, err error) {
	permission := &domain.Permissions{}
	filters := makePermissionIDFilters(input.ID)

	err = impl.repo.Read(ctx, filters, permission)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return out.PermissionsToView(permission), nil
}
