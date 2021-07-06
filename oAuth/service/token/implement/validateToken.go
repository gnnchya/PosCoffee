package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/out"
	"github.com/gnnchya/PosCoffee/oAuth/service/util"
	"github.com/go-oauth2/oauth2/v4/errors"
	"strings"
	"time"
)

func (impl *implementation) ValidateToken(ctx context.Context, input *string) (view *out.ValidateTokenView, err error) {
	header := strings.ReplaceAll(*input, "Bearer ", "")
	token := &domain.OauthToken{}
	filters := impl.filter.MakeAccessTokenFilter(header)
	err = impl.tokenRepository.Read(ctx, filters, token)
	if err != nil {
		return nil, util.Unauthorized(err)
	}

	refresh := time.Unix(token.Expired, 0)
	refreshExpiresIn := time.Hour * 24 * 3
	accessCreateAt := time.Unix(token.Expired, 0)
	accessExpiresIn := time.Hour * 2
	timeNow := time.Now()
	if token.AccessToken != header {
		return nil, util.Unauthorized(errors.ErrInvalidAccessToken)
	} else if token.RefreshToken != "" && token.RefreshExpiresIn != 0 &&
		refresh.Add(refreshExpiresIn).Before(timeNow) {
		return nil, util.Unauthorized(errors.ErrExpiredRefreshToken)
	} else if token.Expired != 0 && accessCreateAt.Add(accessExpiresIn).Before(timeNow) {
		return nil, util.Unauthorized(errors.ErrExpiredAccessToken)
	}

	token.Expired = int64(accessCreateAt.Add(accessExpiresIn).Sub(timeNow).Seconds())

	return out.ValidateTokenToView(token), err
}
