package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type ViewInput struct {
	ID string `json:"id"`
} // @Name StaffCreateInput


func ViewInputToUserDomain(input *ViewInput) (user *domain.ReadStruct) {
	return &domain.ReadStruct{
		ID: input.ID,
	}
}
