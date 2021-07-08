package rolesin

import (
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

type CreateInput struct {
	ID         string    `json:"id" swaggerignore:"true"`
	Name       *Lang     `json:"name" validate:"required"`
	Permission []*string `json:"permissions" validate:"required"`
} // @Name RolesCreateInput

func (input *CreateInput) ToDomain() (roles *domain.Roles) {
	if reflect2.IsNil(input) {
		return &domain.Roles{}
	}

	return &domain.Roles{
		ID:          input.ID,
		Name:        input.Name.ToDomain(),
		Permissions: input.Permission,
		CreatedAt:   carbon.Now().Unix(),
		UpdatedAt:   carbon.Now().Unix(),
	}
}
