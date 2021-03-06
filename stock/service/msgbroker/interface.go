package msgbroker

import "github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"

type Service interface {
	Receiver(topics []msgbrokerin.TopicMsgBroker)
	Sender(topic msgbrokerin.TopicMsgBroker, data interface{}) (err error)
}