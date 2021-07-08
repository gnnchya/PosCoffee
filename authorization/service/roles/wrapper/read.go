package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Read(ctx context.Context, input *rolesin.ReadInput) (view *out.RolesView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.Read")
	defer sp.Finish()

	sp.LogKV("id", input.ID)

	view, err = wrp.service.Read(ctx, input)

	sp.LogKV("view", view)
	sp.LogKV("err", err)

	return view, err
}
