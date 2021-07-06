package grpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/gnnchya/PosCoffee/oAuth/config"
	pb "github.com/gnnchya/PosCoffee/oAuth/service/grpc/protobuf/oauth"
	"github.com/gnnchya/PosCoffee/oAuth/service/token"
	"github.com/gnnchya/PosCoffee/oAuth/service/util"
)

type OauthServiceImpl struct {
	tokenService    token.Service
	filter          util.Filters
	tokenRepository util.Repository
}

const (
	NETWORK = "tcp"
)

func NewServer(conf *config.Config, tokenService token.Service, filter util.Filters, tokenRepository util.Repository) {
	server := OauthServiceImpl{
		tokenService:    tokenService,
		filter:          filter,
		tokenRepository: tokenRepository,
	}

	lis, err := net.Listen(NETWORK, conf.GRPCSenderHost)

	grpcServer := grpc.NewServer()
	pb.RegisterOauthServer(grpcServer, &server)

	// start grpcServer
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
