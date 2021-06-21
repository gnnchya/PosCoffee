package userin

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type DeleteInput struct {
	ID string `bson:"_id" json:"id"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func DeleteInputToUserDomain(input *DeleteInput) (user *domain.DeleteQ) {
	return &domain.DeleteQ{
		ID: input.ID,
		Code: input.Code,
		Err: input.Err,
	}
}

func (input *DeleteInput)DeleteInputToUserDomain() (user *domain.DeleteQ) {
	return &domain.DeleteQ{
		ID:             input.ID,
		Code: input.Code,
		Err: input.Err,
	}
}