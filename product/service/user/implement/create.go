package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}

	//user := userin.CreateInputToUserDomain(input)
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)

	err = impl.repo.Create(ctx, user, user.ID)
	//fmt.Println("output create:", user)

	//if err != nil {
	//	return "", err
	//}

	//if err == impl.sendMsgCreate(input) {
	//	log.Println(err)
	//}

	//time.Sleep(5 * time.Second)
	//_, err = impl.repo.Read(ctx, input.ID)
	if err != nil {
		return "", err
	}

	return ID, nil
}

//func (impl *implementation) sendMsgCreate(input *userin.CreateInput) (err error) {
//	return impl.MsgSender("create", userin.MsgBrokerCreate{
//		Action:         msgbrokerin.ActionCreate,
//		ID:             input.ID,
//		Name:           input.Name,
//		ActualName:     input.ActualName,
//		ActualLastName: input.ActualLastName,
//		Gender:         input.Gender,
//		BirthDate:      input.BirthDate,
//		Height:         input.Height,
//		SuperPower:     input.SuperPower,
//		Alive:          input.Alive,
//		Universe:       input.Universe,
//		Movies:         input.Movies,
//		Enemies:        input.Enemies,
//		FamilyMember:   input.FamilyMember,
//		About:          input.About,
//	})
//}
