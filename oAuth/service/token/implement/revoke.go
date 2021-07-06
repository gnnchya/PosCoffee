package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"strings"

	"github.com/gnnchya/PosCoffee/oAuth/service/util"
)

func (impl *implementation) RevokeToken(ctx context.Context, input *string) (err error) {
	header := strings.ReplaceAll(*input, "Bearer ", "")
	filters := impl.filter.MakeAccessTokenFilter(header)
	token := &domain.OauthToken{}
	err = impl.tokenRepository.Read(ctx, filters, token)
	if err != nil {
		return util.RepoReadErr(err)
	}

	err = impl.tokenRepository.Delete(ctx, filters)
	if err != nil {
		return util.RepoDeleteErr(err)
	}

	return err
}
