package out

import (
	"github.com/gnnchya/PosCoffee/authorize/domain"
)

type PermissionsView struct {
	ID        string `json:"id"`
	Method    string `json:"method"`
	Endpoint  string `json:"endpoint"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
} // @Name PermissionsView

func PermissionsToView(permissions *domain.Permissions) (view *PermissionsView) {
	return &PermissionsView{
		ID:        permissions.ID,
		Method:    permissions.Method,
		Endpoint:  permissions.Endpoint,
		CreatedAt: permissions.CreatedAt,
		UpdatedAt: permissions.UpdatedAt,
	}
}
