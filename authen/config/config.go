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
	GRPCAuthenHost string `env:"GRPC_AUTHEN_HOST" envDefault:"localhost:50057"`


	OauthServerURL   string `env:"OAUTH_SERVER_URL" envDefault:"http://localhost:9090/pos/request"`
	RevokeTokenURL   string `env:"REVOKE_TOKEN_URL" envDefault:"http://localhost:9090/pos/revoke"`
	ValidateTokenURL string `env:"VALIDATE_TOKEN_URL" envDefault:"http://localhost:9090/pos/validate"`
	ClientId         string `env:"CLIENT_ID" envDefault:"49c9a79eJ2ee0S460bH94abx5ea6124231bb"`
	ClientSecret     string `env:"CLIENT_SECRET" envDefault:"28dc28ccJ885fS4861Ha3c0x21619025de7b"`
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
