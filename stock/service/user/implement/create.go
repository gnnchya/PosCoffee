package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Create(ctx context.Context, input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)

	err = impl.repo.Create(ctx, user, user.ID)
	fmt.Println("err in create", err)
	if err != nil {
		return "error create in stock", err
	}

	return user.ID, nil
}

//func (impl *implementation) sendMsgCreate(input *userin.CreateInput) (err error) {
//	return impl.MsgSender("create", userin.MsgBrokerCreate{
//		Action:         msgbrokerin.ActionCreate,
//		ID:             input.ID,
//		ItemName:       input.ItemName,
//		Category:     	input.Category,
//		Amount: 		input.Amount,
//		Unit:         	input.Unit,
//		CostPerUnit:    input.CostPerUnit,
//		EXPDate:        input.EXPDate,
//		ImportDate:     input.ImportDate,
//		Supplier:       input.Supplier,
//		TotalCost:      input.TotalCost,
//		TotalAmount:    input.TotalAmount,
//	})
//}
