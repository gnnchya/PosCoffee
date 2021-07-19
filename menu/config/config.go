package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	// Jaeger config
	JaegerAgentHost string `env:"JAEGER_AGENT_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_AGENT_PORT" envDefault:"6831"`

	//MessageBroker config
	MessageBrokerBackOffTime  int      `env:"MESSAGE_BROKER_BACKOFF_TIME" envDefault:"2"`
	MessageBrokerMaximumRetry int      `env:"MESSAGE_BROKER_MAXIMUM_RETRY" envDefault:"3"`
	MessageBrokerEndpoint     []string `env:"MESSAGE_BROKER_ENDPOINT" envDefault:"localhost:9094"`
	MessageBrokerGroup        string   `env:"MESSAGE_BROKER_GROUP" envDefault:"my-group"`
	MessageBrokerVersion      string   `env:"MESSAGE_BROKER_VERSION" envDefault:"2.6.1"`

	ElasticDBEndpoint     string `env:"ELASTIC_ENDPOINT" envDefault:"http://localhost:9200"`
	ElasticDBUsername     string `env:"ELASTIC_USERNAME" envDefault:"touch"`
	ElasticDBPassword	  string `env:"ELASTIC_PASSWORD" envDefault:"touchja"`

	RedisEndpoint string `env:"REDIS_ENDPOINT" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`

	GRPCAuthenHost string `env:"GRPC_AUTHEN_HOST" envDefault:"localhost:50057"`

}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
