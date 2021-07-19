package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"golang.org/x/crypto/bcrypt"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {

	err = impl.validator.Validate(input)
	if err != nil {
		return "", err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	input.Password = string(hashPassword)

	user := input.ToDomain()
	fmt.Println("Createstruct", user)
	_, err = impl.repo.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

//func (impl *implementation) sendMsgCreate(input *usersin.CreateInput) (err error) {
//	return impl.MsgSender(msgbrokerin.TopicOTP, usersin.MsgBrokerCreate{
//		Action:         msgbrokerin.ActionCreate,
//		ID:             input.ID,
//		Prefix:         input.Prefix,
//		FirstName:      input.FirstName,
//		LastName:       input.LastName,
//		SignUpChannel:  input.SignUpChannel,
//		Email:          input.Email,
//		MobileNumber:   input.MobileNumber,
//		IdentifyType:   input.IdentifyType,
//		IdentifyNumber: input.IdentifyNumber,
//		Password:       input.Password,
//		Passcode:       input.PassCode,
//		MetaData:       input.MetaData,
//		RoleID:         input.RoleID,
//		MemberGroup:    input.MemberGroup,
//	})
//}
