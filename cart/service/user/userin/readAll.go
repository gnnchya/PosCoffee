package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/cart/domain"
)

type ViewAllInput struct {
	PerPage int
	Page    int
} // @Name StaffCreateInput

//func MakeTestViewAllInput() (input *UpdateInput) {
//	return &UpdateInput{
//		ID: "test",
//		// ID:        "test",
//		//Name: "test",
//		// Tel:       "test",
//	}
//}

func ViewAllInputToUserDomain(input *ViewAllInput) (user *domain.ReadAllStruct) {
	return &domain.ReadAllStruct{
		PerPage: input.PerPage,
		Page:    input.Page,
	}
}
