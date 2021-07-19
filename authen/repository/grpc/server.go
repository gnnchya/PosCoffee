package grpc

import (
	"log"
	"net"
)

func (grpcRepo *Config) NetListener() (lis net.Listener, err error) {
	lis, err = net.Listen(grpcRepo.Network, grpcRepo.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return nil, err
	}

	return lis, nil
}
