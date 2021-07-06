package kafka

import (
	"github.com/touchtechnologies-product/message-broker"
	"github.com/touchtechnologies-product/message-broker/common"
)

type Kafka struct {
	message.Broker
}

type Config struct {
	BackOffTime  int
	MaximumRetry int
	Host         []string
	Group        string
	Version      string
}

func New(config *Config) (kafka *Kafka, err error) {
	conf := config.setConfig()
	msgBrokers, err := message.NewBroker(common.KafkaBrokerType, conf)
	if err != nil {
		return nil, err
	}

	kafka = &Kafka{
		msgBrokers,
	}

	return kafka, nil
}

func (conf *Config) setConfig() *common.Config {
	return &common.Config{
		BackOffTime:  conf.BackOffTime,
		MaximumRetry: conf.MaximumRetry,
		Version:      conf.Version,
		Group:        conf.Group,
		Host:         makeHostKafka(conf.Host),
		Debug:        true,
	}
}
