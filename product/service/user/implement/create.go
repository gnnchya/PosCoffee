package implement

import (
	"context"
	"errors"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/calculation"
	"github.com/gnnchya/PosCoffee/product/service/grpcClient/protobuf"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (stock bool, change []domain.ChangeStruct, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return  false, change, err
	}
	fmt.Println("input cart", input.Cart)
	var ingredientList []string
	for _, menu := range input.Cart.Menu{
		for _, ingredient := range menu.Ingredient{
			if found := Find(ingredientList, ingredient.IngredientName); !found{
				ingredientList = append(ingredientList, ingredient.IngredientName)
			}
		}
	}
	inputIngre := &protobuf.RequestToStock{Ingredient: ingredientList}
	res, err := impl.client.SendIngredients(inputIngre)
	fmt.Println("response from stock", res)
	var cost []domain.CalculateCost
	var remainMoney []domain.CreateMoneyStruct
	if res.Stock == true{
		if input.PaymentMethod == "Cash"{
			fmt.Println("hererhehrehrehrherhehre")
			temp, err := impl.repom.ReadMoneyAll(ctx)
			fmt.Println("temp", temp)
			if err != nil{
				return res.Stock,nil , err
			}
			remainMoney, change, err = calculation.Calculation(input.Paid, input.Price, temp)
			fmt.Println("change", change)
			for _,i := range remainMoney{
				err = impl.repom.UpdateByVal(ctx, i, i.Value)
			}
		}
		user := input.CreateInputToUserDomain(cost)
		fmt.Println("user input create:", user)
		err = impl.repo.Create(ctx, user, user.ID)
		if err != nil {
			return res.Stock, change,  err
		}
	}else{
		return res.Stock, change, errors.New(res.Err)
	}

	return res.Stock, change, nil
}

func Find(slice []string, val string) bool{
	for _, item := range slice {
		if item == val {
			return  true
		}
	}
	return false
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
