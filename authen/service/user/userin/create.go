package userin

import (
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"
)

type CreateInput struct {
	ID  			string				`bson:"_id" json:"id"`
	UID  			string				`bson:"uid" json:"uid"`
	Username		string				`bson:"username" json:"username" validate:"required"`
	Password 		string				`bson:"password" json:"password" validate:"required"`
	MetaData		MetaDataStruct	`bson:"meta_data" json:"meta_data"`
	RoleID			[]string		`bson:"role_id" json:"role_id"`
}

func (input *CreateInput) ToDomain() (users *domain.UserStruct) {
	if reflect2.IsNil(input) {
		return &domain.UserStruct{}
	}
	return &domain.UserStruct{
		ID:             input.ID,
		UID:			input.UID,
		Username: 		input.Username,
		Password:       input.Password,
		MetaData:       input.MetaData.ToDomain(),
		RoleID:         input.RoleID,
		CreatedAt:      carbon.Now().Unix(),
		UpdatedAt:      carbon.Now().Unix(),
		DeletedAt: 		0,
	}
}

