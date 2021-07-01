package msgbroker

import "github.com/gnnchya/PosCoffee/authen/service/msgbroker/msgbrokerin"

type Service interface {
	Receiver(topics []msgbrokerin.TopicMsgBroker)
	Sender(topic msgbrokerin.TopicMsgBroker, data interface{}) (err error)
}