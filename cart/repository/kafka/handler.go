package kafka

import (
	"github.com/gnnchya/PosCoffee/cart/service/msgbroker/msgbrokerin"
	"github.com/touchtechnologies-product/message-broker/common"
	)

func (message Kafka) RegisterHandler(topic msgbrokerin.TopicMsgBroker, handler common.Handler) {
	message.Broker.RegisterHandler(string(topic), handler)
}
