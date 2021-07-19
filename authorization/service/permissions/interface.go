package permissions

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/out"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.PermissionsView, err error)
	Create(ctx context.Context, input *permissionsin.CreateInput) (ID string, err error)
	Read(ctx context.Context, input *permissionsin.ReadInput) (user *out.PermissionsView, err error)
	Update(ctx context.Context, input *permissionsin.UpdateInput) (err error)
	Delete(ctx context.Context, input *permissionsin.DeleteInput) (err error)
}
