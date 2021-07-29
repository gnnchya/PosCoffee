package userin

import (
	"github.com/gnnchya/PosCoffee/menu/domain"
)

type ViewAllInput struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
} // @Name StaffCreateInput


func ViewAllInputToUserDomain(input *ViewAllInput) (user *domain.ReadAllStruct) {
	return &domain.ReadAllStruct{
		PerPage: input.PerPage,
		Page:    input.Page,
	}
}
