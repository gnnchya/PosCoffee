package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.PermissionsView, err error) {
	opt.Filters = append(opt.Filters, makeDeletedAtIsNullFilterString())

	total, records, err := impl.repo.List(ctx, opt, &domain.Permissions{})
	if err != nil || len(records) == 0 {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*out.PermissionsView, len(records))
	for i, record := range records {
		items[i] = out.PermissionsToView(record.(*domain.Permissions))
	}

	return total, items, nil
}
