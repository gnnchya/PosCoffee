package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	user := input.UpdateInputToUserDomain()
	err = impl.repo.Update(ctx, input, input.ID)
	//input.Err = err
	//if err == impl.sendMsgUpdate(input){
	//	log.Println(err)
	//}
	//time.Sleep(5 * time.Second)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

//func (impl *implementation) sendMsgUpdate(input *userin.UpdateInput) (err error) {
//	return impl.MsgSender("update", userin.MsgBrokerUpdate{
//		Action:     msgbrokerin.ActionUpdate,
//		ID:        	input.ID,
//		Amount:   	input.Amount,
//		Err : 		input.Err,
//	})
//}
