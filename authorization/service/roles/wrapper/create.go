package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Create(ctx context.Context, input *rolesin.CreateInput) (ID string, err error) {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.Create")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)
	sp.LogKV("Permission", input.Permission)

	ID, err = wrp.service.Create(ctx, input)

	sp.LogKV("ID", ID)
	sp.LogKV("err", err)

	return ID, err
}
