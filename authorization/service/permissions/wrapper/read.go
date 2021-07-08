package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Read(ctx context.Context, input *permissionsin.ReadInput) (view *out.PermissionsView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.permission.Read")
	defer sp.Finish()

	sp.LogKV("id", input.ID)

	view, err = wrp.service.Read(ctx, input)

	sp.LogKV("view", view)
	sp.LogKV("err", err)

	return view, err
}
