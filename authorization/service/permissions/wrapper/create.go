package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Create(ctx context.Context, input *permissionsin.CreateInput) (ID string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.permission.Create")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Method", input.Method)
	sp.LogKV("Endpoint", input.Endpoint)

	ID, err = wrp.service.Create(ctx, input)

	sp.LogKV("ID", ID)
	sp.LogKV("err", err)

	return ID, err
}
