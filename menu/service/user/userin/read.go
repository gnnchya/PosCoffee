package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/menu/domain"
)

type ViewInput struct {
	ID string `json:"id"`
	// Name       string   `son:"name" validate:"required"`
	// ActualName string   `json:"actual_name" validate:"required"`
	// Gender     string   `json:"gender"`
	// BirthDate  int64    `json:"birth_date"`
	// Height     int      `json:"height" validate:"gte=0"`
	// SuperPower []string `json:"super_power"`
	// Alive      bool     `json:"alive"`
} // @Name StaffCreateInput

func MakeTestViewInput() (input *UpdateInput) {
	return &UpdateInput{
		ID: "test",
		// ID:        "test",
		Name: "test",
		// Tel:       "test",
	}
}

func ViewInputToUserDomain(input *ViewInput) (user *domain.ReadStruct) {
	return &domain.ReadStruct{
		ID: input.ID,
	}
}
