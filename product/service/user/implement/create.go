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

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (stock bool, changeAmount []domain.ChangeStruct, changeValue int64, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return  false, changeAmount,changeValue, err
	}
	var ingredientList []string
	for _, menu := range input.Cart.Menu{
		for _, ingredient := range menu.Ingredient{
			if found := Find(ingredientList, ingredient.IngredientName); !found{
				ingredientList = append(ingredientList, ingredient.IngredientName)
			}
		}
	}
	inputIngredient := &protobuf.RequestToStock{Ingredient: ingredientList}
	res, err := impl.client.SendIngredients(inputIngredient)
	var cost []domain.CalculateCost
	var remainMoney []domain.CreateMoneyStruct

	if res.Stock == true{
		if input.PaymentMethod == "Cash"{
			temp, err := impl.repoMoney.ReadMoneyAll(ctx)
			fmt.Println("temp", temp)
			if err != nil{
				return res.Stock,nil ,changeValue, err
			}
			remainMoney, changeAmount,changeValue, err = calculation.Calculation(input.Paid, input.Price, temp)

			for _,i := range remainMoney{
				err = impl.repoMoney.UpdateByVal(ctx, i, i.Value)
			}
		}
		user := input.CreateInputToUserDomain(cost)
		err = impl.repo.Create(ctx, user)
		if err != nil {
			return res.Stock, changeAmount,changeValue,  err
		}
	}else{
		return res.Stock, changeAmount,changeValue, errors.New(res.Err)
	}

	return res.Stock, changeAmount,changeValue, nil
}

func Find(slice []string, val string) bool{
	for _, item := range slice {
		if item == val {
			return  true
		}
	}
	return false
}
