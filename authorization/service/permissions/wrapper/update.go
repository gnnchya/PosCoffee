package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Update(ctx context.Context, input *permissionsin.UpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.permission.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Method", input.Method)
	sp.LogKV("Endpoint", input.Endpoint)

	err = wrp.service.Update(ctx, input)

	sp.LogKV("err", err)

	return err
}
