package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/calculation"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (change map[int64]int64, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return  change, err
	}

	//TODO check with the stock if the ingredients are enough to make
	//TODO input from GRPC from check stock
	// TODO Return from GRPC get err, and []domain.CalculateCost then check if err != nil then return
	// bool, expenses =  []domain.CalculateCost, err
	var cost []domain.CalculateCost
	//if else, if false or err don't calculate change
	var paid int64 // TODO parameter as paid from customer
	var remainMoney []domain.CreateMoneyStruct
	if input.PaymentMethod == "Cash"{
		temp, err := impl.repom.ReadMoneyAll(ctx)
		if err != nil{
			return nil , err
		}
		remainMoney, change, err = calculation.Calculation(paid, input.Price, temp)
		for _,i := range remainMoney{
			err = impl.repom.UpdateByVal(ctx, i, i.Value)
		}
	}
	user := input.CreateInputToUserDomain(cost)
	fmt.Println("user input create:", user)
	err = impl.repo.Create(ctx, user, user.ID)
	if err != nil {
		return change,  err
	}

	return change, nil
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
