package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) List(ctx context.Context, opt *domain.PageOption) (total int, list []*out.PermissionsView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.permission.List")
	defer sp.Finish()

	sp.LogKV("page", opt.Page)
	sp.LogKV("per_page", opt.PerPage)
	sp.LogKV("filters", opt.Filters)

	total, list, err = wrp.service.List(ctx, opt)

	sp.LogKV("total", total)
	sp.LogKV("view", list)
	sp.LogKV("err", err)
	return total, list, err
}
