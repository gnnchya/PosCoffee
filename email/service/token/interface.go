package token

import (
	"context"
	"net/http"

	"github.com/gnnchya/PosCoffee/email/service/token/out"
	"github.com/gnnchya/PosCoffee/email/service/token/tokenin"
)

//go:generate mockery --name=Service
type Service interface {
	Request(ctx context.Context, input *tokenin.RequestInput, r *http.Request) (view *out.TokenView, err error)
	ValidateToken(ctx context.Context, input *string) (view *out.ValidateTokenView, err error)
	Refresh(ctx context.Context, input *tokenin.RefreshInput, r *http.Request) (view *out.TokenView, err error)
	RevokeToken(ctx context.Context, input *string) (err error)
}
