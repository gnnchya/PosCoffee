package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func (impl *implementation) MsgReceiver(ctx context.Context, msg []byte) (err error) {
	msgInput := &userin.MsgBrokerCreate{}
	err = json.Unmarshal(msg, msgInput)
	if err != nil {
		return err
	}

	//msgAuthInput := &
	//err = json.Unmarshal(msg, msgInput)

	fmt.Println(string(msg))
	switch msgInput.Action {
	case msgbrokerin.ActionCreate:
		err = impl.receiveCreateAction(ctx, msgInput)
		if err != nil {
			return err
		}
	case msgbrokerin.ActionUpdate:
		err = impl.receiveUpdateAction(ctx, msgInput)
		if err != nil {
			return err
		}
	case msgbrokerin.ActionDelete:
		err = impl.receiveDeleteAction(ctx, msgInput)
		if err != nil {
			return err
		}
	}

	return
}


func (impl *implementation) receiveCreateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToCreateInput()
	domainUser := input.CreateInputToUserDomain()
	err = impl.repo.Create(ctx, domainUser, domainUser.ID)
	if err != nil {
		return err
	}

	return nil
}

func (impl *implementation) receiveUpdateAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToUpdateInput()
	domainUser := input.UpdateInputToUserDomain()
	err = impl.repo.Update(ctx, domainUser, domainUser.ID)
	if err != nil {
		return err
	}
	return nil
}

func (impl *implementation) receiveDeleteAction(ctx context.Context, msgBrokerInput *userin.MsgBrokerCreate) (err error) {
	input := msgBrokerInput.ToDeleteInput()
	domainUser := input.DeleteInputToUserDomain()
	err = impl.repo.Delete(ctx, domainUser.ID)
	if err != nil {
		return err
	}
	return nil
}