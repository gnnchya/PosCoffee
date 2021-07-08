package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Delete(ctx context.Context, input *permissionsin.DeleteInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.permission.Delete")
	defer sp.Finish()

	sp.LogKV("id", input.ID)

	err = wrp.service.Delete(ctx, input)

	sp.LogKV("err", err)

	return err
}
