package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) Logout(input string) (err error) {
	var ctx context.Context
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.auth.VerifyToken")
	defer sp.Finish()

	sp.LogKV("input", input)

	err = wrp.service.Logout(input)

	sp.LogKV("err", err)

	return err
}
