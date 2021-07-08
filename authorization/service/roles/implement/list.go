package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.RolesView, err error) {
	var roles = &[]out.RolesView{}
	opt.Filters = append(opt.Filters, makeDeletedAtIsNullFilterString())

	pipeline := makePipelineList(opt)

	total, err = impl.repo.Count(ctx, opt.Filters)
	if err != nil || total == 0 {
		return 0, nil, util.RepoListErr(err)
	}

	err = impl.repo.Aggregate(ctx, &pipeline, roles)
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*out.RolesView, len(*roles))
	for i, record := range *roles {
		items[i] = out.RolesToView(&record)
	}

	return total, items, nil
}
