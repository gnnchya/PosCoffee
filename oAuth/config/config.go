package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	// MongoDB config
	MongoDBEndpoint         		string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://touch:touchja@localhost:27014"`
	MongoDBName             		string `env:"MONGODB_NAME" envDefault:"oauth"`
	MongoDBConsumerTableName 		string `env:"MONGODB_CONSUMER_TABLE_NAME" envDefault:"consumer"`
	MongoDBTokenTableName 			string `env:"MONGODB_TOKEN_TABLE_NAME" envDefault:"token"`

	//MessageBroker config
	MessageBrokerBackOffTime  int      `env:"MESSAGE_BROKER_BACKOFF_TIME" envDefault:"2"`
	MessageBrokerMaximumRetry int      `env:"MESSAGE_BROKER_MAXIMUM_RETRY" envDefault:"3"`
	MessageBrokerEndpoint     []string `env:"MESSAGE_BROKER_ENDPOINT" envDefault:"localhost:9094"`
	MessageBrokerGroup        string   `env:"MESSAGE_BROKER_GROUP" envDefault:"my-group"`
	MessageBrokerVersion      string   `env:"MESSAGE_BROKER_VERSION" envDefault:"2.6.1"`

	GRPCSenderHost string `env:"GRPC_SENDER_HOST" envDefault:"localhost:50052"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
