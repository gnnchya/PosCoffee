package roles

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/out"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.RolesView, err error)
	Create(ctx context.Context, input *rolesin.CreateInput) (ID string, err error)
	Read(ctx context.Context, input *rolesin.ReadInput)(role *out.RolesView, err error)
	Update(ctx context.Context, input *rolesin.UpdateInput) (err error)
	Delete(ctx context.Context, input *rolesin.DeleteInput) (err error)
	CheckPermission(ctx context.Context, roles []string, permission string) (allow bool, err error)
}
