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
	GRPCAuthorizeHost string `env:"GRPC_AUTHORIZE_HOST" envDefault:"localhost:50055"`

	OauthServerURL   string `env:"OAUTH_SERVER_URL" envDefault:"http://localhost:9090/pos/request"`
	RevokeTokenURL   string `env:"REVOKE_TOKEN_URL" envDefault:"http://localhost:9090/pos/revoke"`
	ValidateTokenURL string `env:"VALIDATE_TOKEN_URL" envDefault:"http://localhost:9090/pos/validate"`
	ClientId         string `env:"CLIENT_ID" envDefault:"52f35122J4625S45b9Hb8e8x65d270f9048a"`
	ClientSecret     string `env:"CLIENT_SECRET" envDefault:"6df38733J5791S4a52Hbe76xa278f8a36cbb"`
	RedirectUri      string `env:"REDIRECT_URI" envDefault:"localhost"`
	GrantType        string `env:"GRANT_TYPE" envDefault:"implement"`

	MongoDBEndpoint         string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://touch:touchja@localhost:27016"`
	MongoDBName             string `env:"MONGODB_NAME" envDefault:"user"`
	MongoDBTableName 		string `env:"MONGODB_TABLE_NAME" envDefault:"user"`

	VerifyUrl 		string		`env:"VERIFY_URL" envDefault:"http://localhost:8085/pos/verify"`
	Email 		string		`env:"EMAIL" envDefault:"62011155@kmitl.ac.th"`
	EmailPassword 		string		`env:"EMAIL_PASSWORD" envDefault:"agjxdrywfdtueorp"`

}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
