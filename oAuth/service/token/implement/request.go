package implement

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/server"

	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/out"
	"github.com/gnnchya/PosCoffee/oAuth/service/token/tokenin"
)

func (impl *implementation) Request(ctx context.Context, input *tokenin.RequestInput, r *http.Request) (view *out.TokenView, err error) {
	consumer := &domain.ConsumerStruct{}
	filters := makeConsumerFilters(impl.filter, input.ClientID, input.ClientSecret)

	err = impl.consumerRepository.Read(ctx, filters, consumer)
	if err != nil {
		return nil, err
	}

	checkAccessToken := &domain.TokenStruct{}
	filters = makeUserIDFilters(impl.filter, input.UID)

	err = impl.consumerRepository.Read(ctx, filters, checkAccessToken)
	if err != nil {
		return nil, err
	}

	checkToken := &domain.TokenStruct{}
	pastAccessToken := "Bearer "+checkAccessToken.AccessToken
	fmt.Println( "past access token", pastAccessToken)
	filters = makeUserIDFilters(impl.filter, input.UID)
	err = impl.tokenRepository.Read(ctx, filters, checkToken)
	if err == nil {
		impl.RevokeToken(ctx, &pastAccessToken)
	}

	impl.oauthRepository.ClientStore(consumer.ClientID, consumer.ClientSecret, consumer.RedirectUri)

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
	input.CreatedAt = ti.GetAccessCreateAt().Unix()
	input.AccessToken = ti.GetAccess()
	input.RefreshToken = ti.GetRefresh()
	input.AccessExpire = int64(accessExpiresIn.Seconds())
	input.RefreshExpire = int64(refreshExpiresIn.Seconds())

	token := input.ToDomain()

	err = impl.tokenRepository.Create(ctx, token)
	if err != nil {
		return nil, err
	}

	tokenView := &domain.TokenStruct{
		AccessToken:  input.AccessToken,
		RefreshToken: input.RefreshToken,
		AccessExpire: int64(ti.GetAccessExpiresIn()),
	}

	return out.TokenToView(tokenView), err
}
