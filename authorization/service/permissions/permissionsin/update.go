package permissionsin

import (
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

type UpdateInput struct {
	ID       string `json:"id" validate:"required" swaggerignore:"true"`
	Method   string `json:"method" validate:"required,oneof=GET POST PUT DELETE"`
	Endpoint string `json:"endpoint" validate:"required"`
} // @Name PermissionsUpdateInput

func (input *UpdateInput) ToDomain() (permissions *domain.Permissions) {
	if reflect2.IsNil(input) {
		return &domain.Permissions{}
	}

	return &domain.Permissions{
		ID:        input.ID,
		Method:    input.Method,
		Endpoint:  input.Endpoint,
		UpdatedAt: carbon.Now().Unix(),
	}
}
