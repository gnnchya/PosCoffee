package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) CheckPermission(ctx context.Context, roles []string, permission string) (allow bool, err error) {
	var out out.RolesView
	fmt.Println("permission", permission)
	fmt.Println("roles:", roles)
	if roles == nil{
		return false, nil
	}
	for _, role := range roles{
		filter := makeRoleIDFilters(role)
		err = impl.repo.Read(ctx, filter, out)
		fmt.Println("err", err)
		if err != nil {
			return false, util.RepoReadErr(err)
		}
		for _, per := range out.Permissions{
			fmt.Println("method+url", per.Method+per.Endpoint)
			if per.Method+per.Endpoint == permission{
				return true, nil
			}
		}
	}

	return true, nil
}