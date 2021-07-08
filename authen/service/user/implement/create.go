package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"log"

	"github.com/modern-go/reflect2"
	"golang.org/x/crypto/bcrypt"

)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	defer func() {
		if !reflect2.IsNil(err) {
			return
		}

		if err = impl.sendMsgCreate(input); err != nil {
			log.Println(err)
		}
	}()

	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	//encodeIdentifyNumber, err := util.Encrypt([]byte(input.IdentifyNumber), impl.cryptPassPhrase)
	//if err != nil {
	//	return "", util.RepoCreateErr(err)
	//}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	//hashPassCode, err := bcrypt.GenerateFromPassword([]byte(input.PassCode), bcrypt.DefaultCost)
	//if err != nil {
	//	return "", util.RepoCreateErr(err)
	//}

	//input.ID = impl.uuid.Generate()
	//input.IdentifyNumber = encodeIdentifyNumber
	input.Password = string(hashPassword)
	//input.PassCode = string(hashPassCode)

	user := input.ToDomain()

	_, err = impl.repo.Create(ctx, user)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return user.ID, nil
}

func (impl *implementation) sendMsgCreate(input *usersin.CreateInput) (err error) {
	return impl.MsgSender(msgbrokerin.TopicOTP, usersin.MsgBrokerCreate{
		Action:         msgbrokerin.ActionCreate,
		ID:             input.ID,
		Prefix:         input.Prefix,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		SignUpChannel:  input.SignUpChannel,
		Email:          input.Email,
		MobileNumber:   input.MobileNumber,
		IdentifyType:   input.IdentifyType,
		IdentifyNumber: input.IdentifyNumber,
		Password:       input.Password,
		Passcode:       input.PassCode,
		MetaData:       input.MetaData,
		RoleID:         input.RoleID,
		MemberGroup:    input.MemberGroup,
	})
}
