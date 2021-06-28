package grpc

import (
	"log"
	"net"
)

func (grpc *Config) NetListener() (lis net.Listener, err error) {
	lis, err = net.Listen(grpc.Network, grpc.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return nil, err
	}

	return lis, nil
}
