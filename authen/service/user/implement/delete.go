package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	fmt.Println("user input delete:", input)

	//TODO DElete duay
	//fmt.Println("output del:", user)
	//fmt.Println("err del:", err)
	//if err == impl.sendMsgDelete(input){
	//	log.Println(err)
	//}
	//time.Sleep(5 * time.Second)
	//_, err = impl.repo.View(ctx, input.ID)
	if err != nil {
		return "", err
	}
	return input.ID, err
}

//func (impl *implementation) sendMsgDelete(input *userin.DeleteInput) (err error) {
//	return impl.MsgSender("delete", userin.MsgBrokerDelete{
//		Action:     msgbrokerin.ActionDelete,
//		ID:             input.ID,
//	})
//}

