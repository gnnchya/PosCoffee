package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"
)

type UpdateInput struct {
	ID             string       `json:"id"`
	Prefix         *Lang 		`json:"prefix"`
	FirstName      *Lang 		`json:"first_name" validate:"required"`
	LastName       *Lang 		`json:"last_name" validate:"required"`
	SignUpChannel  string       `json:"sign_up_channel"`
	Email          string       `json:"email"`
	MobileNumber   string       `json:"mobile_number" validate:"required"`
	IdentifyType   string       `json:"identify_type" validate:"required"`
	IdentifyNumber string       `json:"identify_number" validate:"required"`
	Password       string       `json:"password" validate:"required"`
	PassCode       string       `json:"passcode"`
	RoleID         []string     `json:"role_id"`
	MemberGroup    *MemberGroup `json:"member_group"`
	CreatedAt      int64        `json:"-"`
	*MetaData      				`validate:"required"`
}

func (input *UpdateInput)UpdateInputToUserDomain() (user *domain.Users) {
	if reflect2.IsNil(input) {
		return &domain.Users{}
	}

	return &domain.Users{
		ID:             input.ID,
		Prefix:         input.Prefix.ToDomain(),
		FirstName:      input.FirstName.ToDomain(),
		LastName:       input.LastName.ToDomain(),
		SignUpChannel:  input.SignUpChannel,
		Email:          input.Email,
		MobileNumber:   input.MobileNumber,
		IdentifyType:   input.IdentifyType,
		IdentifyNumber: input.IdentifyNumber,
		Password:       input.Password,
		PassCode:       input.PassCode,
		RoleID:         input.RoleID,
		MemberGroup:    input.MemberGroup.ToDomain(),
		MetaData:       input.MetaData.ToDomain(),
		CreatedAt:      input.CreatedAt,
		UpdatedAt:      carbon.Now().Unix(),
	}
}