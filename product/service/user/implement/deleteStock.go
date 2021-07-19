package implement

import (
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"log"
)

func (impl *implementation) DeleteStock(input *userin.DeleteStockInput) (ID string, err error) {
	if err == impl.sendMsgDelete(input){
		log.Println(err)
	}

	if err != nil {
		return "", err
	}
	return ID, nil
}

func (impl *implementation) sendMsgDelete(input *userin.DeleteStockInput) (err error) {
	return impl.MsgSender("delete", userin.MsgBrokerDelete{
		Action:     msgbrokerin.ActionDelete,
		ID:             input.ID,
	})
}

