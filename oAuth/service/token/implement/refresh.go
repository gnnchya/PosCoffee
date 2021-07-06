package implement

import (
	"context"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/server"

	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/out"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/tokenin"
	"github.com/gnnchya/PosCoffee/oAuth/service/util"
)

func (impl *implementation) Refresh(ctx context.Context, input *tokenin.RefreshInput, r *http.Request) (view *out.TokenView, err error) {
	consumer := &domain.Consumer{}
	filters := makeClientFilters(impl.filter, input.ClientID, input.ClientSecret)

	err = impl.consumerRepository.Read(ctx, filters, consumer)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	impl.oauthRepository.ClientStore(consumer.ClientId, consumer.ClientSecret, consumer.RedirectUri)

	srv := impl.oauthRepository.NewServer()
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	gt, tgr, err := srv.ValidationTokenRequest(r)
	if err != nil {
		return nil, err
	}

	ti, err := srv.GetAccessToken(ctx, gt, tgr)
	if err != nil {
		return nil, err
	}

	accessExpiresIn := time.Hour * 2
	refreshExpiresIn := time.Hour * 24 * 3

	ti.SetAccessExpiresIn(accessExpiresIn)
	ti.SetRefreshExpiresIn(refreshExpiresIn)

	token := &domain.OauthToken{
		ID:               impl.uuid.Generate(),
		UserId:           input.UserId,
		AccessToken:      ti.GetAccess(),
		RefreshToken:     ti.GetRefresh(),
		Expired:          ti.GetAccessCreateAt().Unix(),
		ExpiryDate:       int64(accessExpiresIn.Seconds()),
		RefreshExpiresIn: int64(refreshExpiresIn.Seconds()),
	}

	_, err = impl.tokenRepository.Create(ctx, token)
	if err != nil {
		return nil, util.RepoCreateErr(err)
	}

	return out.TokenToView(token), err
}
