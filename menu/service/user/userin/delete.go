package userin

import (
	"github.com/gnnchya/PosCoffee/menu/domain"
)

type DeleteInput struct {
	ID string `bson:"_id" json:"id"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func DeleteInputToUserDomain(input *DeleteInput) (user *domain.DeleteStruct) {
	return &domain.DeleteStruct{
		ID: input.ID,
		Code: input.Code,
		Err: input.Err,
	}
}

func (input *DeleteInput)DeleteInputToUserDomain() (user *domain.DeleteStruct) {
	return &domain.DeleteStruct{
		ID:             input.ID,
		Code: input.Code,
		Err: input.Err,
	}
}