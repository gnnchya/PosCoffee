package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *rolesin.ReadInput) (view *out.RolesView, err error) {
	var role = make([]out.RolesView, 1)
	var aggregateData out.RolesView
	filters := makeRoleIDFilters(input.ID)

	pipeline := makePipelineRead(&filters)
	err = impl.repo.Aggregate(ctx, &pipeline, &role)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	if len(role) == 0 {
		return nil, util.RepoReadErr(err)
	}

	aggregateData = role[0]

	return out.RolesToView(&aggregateData), nil
}