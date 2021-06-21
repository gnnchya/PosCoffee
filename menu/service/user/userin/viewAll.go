package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type ViewAllInput struct {
	PerPage int
	Page    int
} // @Name StaffCreateInput

func MakeTestViewAllInput() (input *UpdateInput) {
	return &UpdateInput{
		ID: "test",
		// ID:        "test",
		Name: "test",
		// Tel:       "test",
	}
}

func ViewAllInputToUserDomain(input *ViewAllInput) (user *domain.ViewByPageQ) {
	return &domain.ViewByPageQ{
		PerPage: input.PerPage,
		Page:    input.Page,
	}
}
