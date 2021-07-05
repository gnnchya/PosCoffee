package implement

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"log"
)

func (impl *implementation) CreateStock(input *userin.CreateStockInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "validate error", err
	}

	user := input.CreateStockInputToUserDomain()

	if err == impl.sendMsgCreate(user) {
		log.Println(err)
	}

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
