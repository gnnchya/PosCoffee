package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Update(ctx context.Context, input *rolesin.UpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)
	sp.LogKV("Permission", input.Permission)

	err = wrp.service.Update(ctx, input)

	sp.LogKV("err", err)

	return err
}
