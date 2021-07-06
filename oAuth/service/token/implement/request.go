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

func (impl *implementation) Request(ctx context.Context, input *tokenin.RequestInput, r *http.Request) (view *out.TokenView, err error) {

	consumer := &domain.Consumer{}
	filters := makeConsumerFilters(impl.filter, input.ClientID, input.ClientSecret, input.RedirectUri)

	err = impl.consumerRepository.Read(ctx, filters, consumer)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	impl.oauthRepository.ClientStore(consumer.ClientId, consumer.ClientSecret, consumer.RedirectUri)

	srv := impl.oauthRepository.NewServer()
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		userID = username
		return
	})

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

	ti.SetAccessExpiresIn(accessExpiresIn / time.Second)
	ti.SetRefreshExpiresIn(refreshExpiresIn / time.Second)

	input.ID = impl.uuid.Generate()
	input.Expired = ti.GetAccessCreateAt().Unix()
	input.AccessToken = ti.GetAccess()
	input.RefreshToken = ti.GetRefresh()
	input.ExpiryDate = int64(accessExpiresIn.Seconds())
	input.RefreshExpiresIn = int64(refreshExpiresIn.Seconds())

	token := input.ToDomain()

	_, err = impl.tokenRepository.Create(ctx, token)
	if err != nil {
		return nil, util.RepoCreateErr(err)
	}

	tokenView := &domain.OauthToken{
		AccessToken:  input.AccessToken,
		RefreshToken: input.RefreshToken,
		Expired:      int64(ti.GetAccessExpiresIn()),
	}

	return out.TokenToView(tokenView), err
}
