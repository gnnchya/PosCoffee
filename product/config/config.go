package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	// MongoDB config
	MongoDBEndpoint         string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://touch:touchja@localhost:27018"`
	MongoDBName             string `env:"MONGODB_NAME" envDefault:"product"`
	MongoDBTableName 		string `env:"MONGODB_TABLE_NAME" envDefault:"transactions"`

	MongoDBEndpointMoney    string `env:"MONGODB_MONEY_ENDPOINT" envDefault:"mongodb://touch:touchja@localhost:27018"`
	MongoDBNameMoney        string `env:"MONGODB_MONEY_NAME" envDefault:"product"`
	MongoDBTableNameMoney  	string `env:"MONGODB_MONEY_TABLE_NAME" envDefault:"money"`


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
	// gRPC config
	GRPCHost       string `env:"GRPC_HOST" envDefault:"localhost:50051"`
	GRPCSenderHost string `env:"GRPC_SENDER_HOST" envDefault:"localhost:50052"`
	GRPCSenderReportHost string `env:"GRPC_SENDER_REPORT_HOST" envDefault:"localhost:50053"`
	GRPCMenu string `env:"GRPC_SENDER_HOST" envDefault:"localhost:50054"`

}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
