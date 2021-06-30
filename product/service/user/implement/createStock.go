package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"log"
)

func (impl *implementation) CreateStock(ctx context.Context, input *userin.CreateStockInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}

	user := input.CreateStockInputToUserDomain()
	fmt.Println("user input create:", user)

	if err == impl.sendMsgCreate(user) {
		log.Println(err)
		return "", err
	}

	//time.Sleep(5 * time.Second)
	//_, err = impl.repo.Read(ctx, input.ID)
	if err != nil {
		return "", err
	}

	return ID, nil
}

func (impl *implementation) sendMsgCreate(input *domain.CreateStockStruct) (err error) {
	return impl.MsgSender("create", userin.MsgBrokerCreate{
		Action:         msgbrokerin.ActionCreate,
		ID:             input.ID,
		ItemName:       input.ItemName,
		Category:     	input.Category,
		Amount: 		input.Amount,
		Unit:         	input.Unit,
		CostPerUnit:    input.CostPerUnit,
		EXPDate:        input.EXPDate,
		ImportDate:     input.ImportDate,
		Supplier:       input.Supplier,
		TotalCost:      input.TotalCost,
		TotalAmount:    input.TotalAmount,
	})
}
