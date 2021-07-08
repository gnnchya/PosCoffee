package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Delete(ctx context.Context, input *rolesin.DeleteInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.Delete")
	defer sp.Finish()

	sp.LogKV("id", input.ID)

	err = wrp.service.Delete(ctx, input)

	sp.LogKV("err", err)

	return err
}
