package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/product/domain"
)

type ReadAllInput struct {
	PerPage int
	Page    int
}

func ReadAllInputToUserDomain(input *ReadAllInput) (user *domain.ReadOrderByPageStruct) {
	return &domain.ReadOrderByPageStruct{
		PerPage: input.PerPage,
		Page:    input.Page,
	}
}

type ReadCategoryAllInput struct {
	Category 	string 	`bson:"category" json:"category"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
}

type ReadNameAllInput struct {
	ItemName 	string 	`bson:"item_name" json:"item_name"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
}
