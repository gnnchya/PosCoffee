package permissionsin

import (
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

type CreateInput struct {
	ID       string `json:"id"`
	Method   string `json:"method" validate:"required,oneof=GET POST PUT DELETE"`
	Endpoint string `json:"endpoint" validate:"required"`
} // @Name PermissionsCreateInput

func (input *CreateInput) ToDomain() (permissions *domain.Permissions) {
	if reflect2.IsNil(input) {
		return &domain.Permissions{}
	}

	return &domain.Permissions{
		ID:        input.ID,
		Method:    input.Method,
		Endpoint:  input.Endpoint,
		CreatedAt: carbon.Now().Unix(),
		UpdatedAt: carbon.Now().Unix(),
	}
}
