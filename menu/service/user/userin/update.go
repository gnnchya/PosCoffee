package userin

import (
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

type UpdateInput struct {
	ID         string   `json:"id"`
	Name       string   `json:"name" validate:"required"`
	ActualName string   `json:"actual_name" validate:"required"`
	ActualLastName string `json:"actual_last_name"`
	Gender     string   `json:"gender"`
	BirthDate  int64    `json:"birth_date"`
	Height     int      `json:"height" validate:"gte=0"`
	SuperPower []string `json:"super_power"`
	Alive      bool     `json:"alive"`
	Universe       string   `json:"universe"`
	Movies         []string `json:"movies"`
	Enemies        []string `json:"enemies"`
	FamilyMember   []string `json:"family_member"`
	About          string   `json:"about"`
	Code int `json:"code"`
	Err error `json:"err"`
}

func UpdateInputToUserDomain(input *UpdateInput) (user *domain.UpdateQ) {
	return &domain.UpdateQ{
		ID:             input.ID,
		Name:           input.Name,
		ActualName:     input.ActualName,
		ActualLastName: input.ActualLastName,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		Height:         input.Height,
		SuperPower:     input.SuperPower,
		Alive:          input.Alive,
		Universe:       input.Universe,
		Movies:         input.Movies,
		Enemies:        input.Enemies,
		FamilyMember:   input.FamilyMember,
		About:          input.About,
		Code: input.Code,
		Err: input.Err,
	}
}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.UpdateQ) {
	return &domain.UpdateQ{
		ID:             input.ID,
		Name:           input.Name,
		ActualName:     input.ActualName,
		ActualLastName: input.ActualLastName,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		Height:         input.Height,
		SuperPower:     input.SuperPower,
		Alive:          input.Alive,
		Universe:       input.Universe,
		Movies:         input.Movies,
		Enemies:        input.Enemies,
		FamilyMember:   input.FamilyMember,
		About:          input.About,
		Code: input.Code,
		Err: input.Err,
	}
}