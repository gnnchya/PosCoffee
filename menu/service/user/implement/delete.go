package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/menu/service/msgbroker/msgbrokerin"
	"log"
	"time"

	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	//defer func(){
	//	if !reflect2.IsNil(err){
	//		return
	//	}
	//	if err == impl.sendMsgDelete(input){
	//		log.Println(err)
	//	}
	//}()
	user := userin.DeleteInputToUserDomain(input)
	fmt.Println("user input delete:", user)

	//err = impl.repo.Delete(ctx, user.ID)
	//fmt.Println("output del:", user)
	//fmt.Println("err del:", err)
	if err == impl.sendMsgDelete(input){
		log.Println(err)
	}
	time.Sleep(5 * time.Second)
	_, err = impl.repo.View(ctx, input.ID)
	if err != nil {
		return "", err
	}
	return user.ID, err
}

func (impl *implementation) sendMsgDelete(input *userin.DeleteInput) (err error) {
	return impl.MsgSender("delete", userin.MsgBrokerDelete{
		Action:     msgbrokerin.ActionDelete,
		ID:             input.ID,
	})
}

