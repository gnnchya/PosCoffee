package userin

import (
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type ViewInput struct {
	CustomerID string `json:"customer_id"`
} // @Name StaffCreateInput


func ViewInputToUserDomain(input *ViewInput) (user *domain.ReadStruct) {
	return &domain.ReadStruct{
		CustomerID: input.CustomerID,
	}
}
