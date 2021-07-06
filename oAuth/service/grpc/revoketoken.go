package grpc

import (
	"context"

	"github.com/gnnchya/PosCoffee/oAuth/domain"
	pb "github.com/gnnchya/PosCoffee/oAuth/service/grpc/protobuf/oauth"
)

func (impl *OauthServiceImpl) RevokeToken(ctx context.Context, data *pb.RevokeRequest) (*pb.RevokeReply, error) {
	resp := &pb.RevokeReply{}

	filters := impl.filter.MakeUserIDString(data.UserID)
	token := &domain.OauthToken{}
	err := impl.tokenRepository.Read(ctx, filters, token)
	if err != nil {
		return resp, nil
	}

	err = impl.tokenRepository.DeleteMany(ctx, filters)
	if err != nil {
		return resp, nil
	}

	resp.Status = "success"

	return resp, nil
}
