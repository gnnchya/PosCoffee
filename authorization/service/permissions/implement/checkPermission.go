package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) CheckPermission(ctx context.Context, roles []string, permission string) (allow bool, err error) {
	var role = make([]out.RolesView, 1)
	var aggregateData out.RolesView
	filters := makeRoleIDFilters(input.ID)

	pipeline := makePipelineRead(&filters)
	err = impl.repo.Aggregate(ctx, &pipeline, &role)
	//TODO check if the permission in in the list of roles
	if err != nil {
		return false, util.RepoCreateErr(err)
	}

	return true, nil
}