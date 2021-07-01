package wrapper

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/autenitcationin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/out"
	"github.com/opentracing/opentracing-go"
)

func (wrp wrapper) GenerateToken(input *autenitcationin.LoginInput) (token *out.Token, err error) {
	var ctx context.Context
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.auth.GenerateToken")
	defer sp.Finish()

	sp.LogKV("Username", input.Username)
	sp.LogKV("Password", input.Password)

	token, err = wrp.service.GenerateToken(input)

	sp.LogKV("token", token)
	sp.LogKV("err", err)

	return token, err
}
