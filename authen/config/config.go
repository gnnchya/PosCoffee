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

	GRPCHost       string `env:"GRPC_HOST" envDefault:"localhost:50051"`
	GRPCSenderHost string `env:"GRPC_SENDER_HOST" envDefault:"localhost:50052"`

	OauthServerURL   string `env:"OAUTH_SERVER_URL" envDefault:"http://localhost:9090/api/v1/request"`
	RevokeTokenURL   string `env:"REVOKE_TOKEN_URL" envDefault:"http://localhost:9090/api/v1/revoke"`
	ValidateTokenURL string `env:"VALIDATE_TOKEN_URL" envDefault:"http://localhost:9090/api/v1/validate"`
	ClientId         string `env:"CLIENT_ID" envDefault:"948cb005Jed97S4726Hbc20xc5dda3384a82"`
	ClientSecret     string `env:"CLIENT_SECRET" envDefault:"55ba3288J7fdaS483bH8f1fxacef36455013"`
	RedirectUri      string `env:"REDIRECT_URI" envDefault:"localhost"`
	GrantType        string `env:"GRANT_TYPE" envDefault:"password"`

	MongoDBEndpoint         string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://touch:touchja@localhost:27016"`
	MongoDBName             string `env:"MONGODB_NAME" envDefault:"user"`
	MongoDBTableName 		string `env:"MONGODB_TABLE_NAME" envDefault:"user"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
