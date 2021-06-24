package grpc

type Config struct {
	Network string
	Port    string
}

func New(config *Config) (grpc *Config) {
	return &Config{config.Network, config.Port}
}
