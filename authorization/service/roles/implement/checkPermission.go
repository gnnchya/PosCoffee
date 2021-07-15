package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) CheckPermission(ctx context.Context, roles []string, permission string) (allow bool, err error) {
	var out = make([]out.RolesView, 1)
	fmt.Println("permission", permission)
	fmt.Println("roles:", roles)
	if roles == nil{
		return false, nil
	}

	for _, role := range roles{
		filter := makeRoleIDFilters(role)
		pipeline := makePipelineRead(&filter)
		err = impl.repo.Aggregate(ctx, &pipeline, &out)
		fmt.Println("out", out)
		fmt.Println("err", err)
		if err != nil {
			return false, util.RepoReadErr(err)
		}
		for _, pers := range out{
			for _, per := range pers.Permissions{
				fmt.Println("method+url", per.Method+per.Endpoint)
				if per.Method+per.Endpoint == permission{
					return true, nil
				}
			}

		}
	}

	return false, nil
}