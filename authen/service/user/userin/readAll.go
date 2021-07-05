package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
)

type ViewAllInput struct {
	PerPage int
	Page    int
} // @Name StaffCreateInput

func ViewAllInputToUserDomain(input *ViewAllInput) (user *domain.ReadAllStruct) {
	return &domain.ReadAllStruct{
		PerPage: input.PerPage,
		Page:    input.Page,
	}
}
