package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	// MongoDB config
	MongoDBEndpoint            string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://touch:touchja@localhost:27015"`
	MongoDBName                string `env:"MONGODB_NAME" envDefault:"authorization_service"`
	MongoDBRoleTableName       string `env:"MONGODB_ROLE_TABLE_NAME" envDefault:"roles"`
	MongoDBPermissionTableName string `env:"MONGODB_PERMISSION_TABLE_NAME" envDefault:"permissions"`

	// Jaeger config
	JaegerAgentHost string `env:"JAEGER_AGENT_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_AGENT_PORT" envDefault:"6831"`

	// Swagger config
	SwaggerHost string `env:"SWAGGER_HOST" envDefault:"localhost:8080"`
	BasePath    string `env:"BASE_PATHS" envDefault:"/api/v1"`

	// gRPC config
	GRPCHost       string `env:"GRPC_HOST" envDefault:"localhost:50051"`
	GRPCSenderHost string `env:"GRPC_SENDER_HOST" envDefault:"localhost:50052"`
	GRPCAuthorizeHost string `env:"GRPC_AUTHORIZE_HOST" envDefault:"localhost:50055"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
