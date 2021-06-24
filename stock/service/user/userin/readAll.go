package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

type ReadCategoryAllInput struct {
	Category 	string 	`bson:"category" json:"category"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
	Code 		int 	`json:"code"`
	Err 		error 	`json:"err"`
}


func ReadCategoryInputToUserDomain(input *ReadCategoryAllInput) (user *domain.ReadCategoryByPageStruct) {
	return &domain.ReadCategoryByPageStruct{
		Category: 	input.Category,
		PerPage: 	input.PerPage,
		Page:    	input.Page,
		Code: 		input.Code,
		Err: 		input.Err,
	}
}

type ReadNameAllInput struct {
	ItemName 	string 	`bson:"item_name" json:"item_name"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
	Code 		int 	`json:"code"`
	Err 		error 	`json:"err"`
}

func ReadNameInputToUserDomain(input *ReadNameAllInput) (user *domain.ReadNameByPageStruct) {
	return &domain.ReadNameByPageStruct{
		ItemName: 	input.ItemName,
		PerPage: 	input.PerPage,
		Page:    	input.Page,
		Code: 		input.Code,
		Err: 		input.Err,
	}
}