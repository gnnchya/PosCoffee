package implement

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

const InvalidInputTypeErr string = "invalid authentication type"

func (impl *implementation) MsgSender(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
	fmt.Println("enter msgsender", topic)
	switch topic {
	case msgbrokerin.TopicCreate:
		err = impl.senderCreate(topic, input)
		if err != nil {
			return err
		}
	case msgbrokerin.TopicUpdate:
		err = impl.senderUpdate(topic, input)
		if err != nil {
			return err
		}
	case msgbrokerin.TopicDelete:
		err = impl.senderDelete(topic, input)
		if err != nil {
			return err
		}
	}
	return
}

func (impl *implementation) senderCreate(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
	create, ok := input.(userin.MsgBrokerCreate) //set data that will be send to kafka
	if !ok {
		return errors.New(InvalidInputTypeErr)
	}

	msg, err := json.Marshal(create)
	if err != nil {
		return err
	}

	err = impl.mBroker.Producer(topic, msg)
	fmt.Println("producer", err)
	if err != nil {
		return err
	}

	return
}

func (impl *implementation) senderUpdate(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
	update, ok := input.(userin.MsgBrokerUpdate) //set data that will be send to kafka
	if !ok {
		return errors.New(InvalidInputTypeErr)
	}

	msg, err := json.Marshal(update)
	if err != nil {
		return err
	}

	err = impl.mBroker.Producer(topic, msg)
	fmt.Println("producer", err)
	if err != nil {
		return err
	}

	return
}

func (impl *implementation) senderDelete(topic msgbrokerin.TopicMsgBroker, input interface{}) (err error) {
	del, ok := input.(userin.MsgBrokerDelete) //set data that will be send to kafka
	if !ok {
		return errors.New(InvalidInputTypeErr)
	}

	msg, err := json.Marshal(del)
	if err != nil {
		return err
	}

	err = impl.mBroker.Producer(topic, msg)
	fmt.Println("producer", err)
	if err != nil {
		return err
	}

	return
}