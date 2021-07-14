package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) CheckPermission(ctx context.Context, roles []string, permission string) (allow bool, err error) {
	var out out.RolesView
	for _, role := range roles{
		filter := makeRoleIDFilters(role)
		err = impl.repo.Read(ctx, filter, out)
		if err != nil {
			return false, util.RepoReadErr(err)
		}
		for _, per := range out.Permissions{
			if per.Method+per.Endpoint == permission{
				return true, nil
			}
		}
	}

	return true, nil
}