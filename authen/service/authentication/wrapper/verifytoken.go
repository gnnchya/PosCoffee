package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) VerifyToken(accessToken string) (userID *string, err error) {
	var ctx context.Context
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.auth.VerifyToken")
	defer sp.Finish()

	sp.LogKV("accessToken", accessToken)

	userID, err = wrp.service.VerifyToken(accessToken)

	sp.LogKV("userID", userID)
	sp.LogKV("err", err)

	return userID, err
}
