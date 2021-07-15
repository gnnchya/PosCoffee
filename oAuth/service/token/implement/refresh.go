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

func (impl *implementation) Refresh(ctx context.Context, input *tokenin.RefreshInput, r *http.Request) (view *out.TokenView, err error) {
	consumer := &domain.ConsumerStruct{}
	filters := makeClientFilters(impl.filter, input.ClientID, input.ClientSecret)

	err = impl.consumerRepository.Read(ctx, filters, consumer)
	if err != nil {
		return nil, err
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
	tgr.UserID = input.UID
	tgr.Scope = "all"
	fmt.Println("refresh tgr", tgr.Refresh)
	//srv.UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (string, error) {
	//	fmt.Println("here1")
	//	return "", errors.ErrAccessDenied
	//}
	//
	//srv.PasswordAuthorizationHandler = func(username, password string) (string, error) {
	//	fmt.Println("here2")
	//	return "", errors.ErrAccessDenied
	//}
	fmt.Println("validation token request refresh", err)
	if err != nil {
		return nil, err
	}
	fmt.Println("gt", gt)
	fmt.Println("tgr", tgr)
	ti, err := srv.GetAccessToken(ctx, gt, tgr)
	fmt.Println("GetAccess err", err)
	if err != nil {
		return nil, err
	}

	accessExpiresIn := time.Hour * 2
	//refreshExpiresIn := time.Hour * 24 * 3

	ti.SetAccessExpiresIn(accessExpiresIn)
	//ti.SetRefreshExpiresIn(refreshExpiresIn)

	token := &domain.TokenStruct{
		ID:               impl.uuid.Generate(),
		UID:           	  input.UID,
		AccessToken:      ti.GetAccess(),
		RefreshToken:     ti.GetRefresh(),
		AccessExpire:     int64(accessExpiresIn.Seconds()),
		RefreshExpire: 	  int64(ti.GetRefreshExpiresIn().Seconds()),
		CreatedAt:		  ti.GetAccessCreateAt().Unix(),
	}

	err = impl.tokenRepository.Create(ctx, token)
	fmt.Println("create err refresh", err)
	if err != nil {
		return nil, err
	}

	return out.TokenToView(token), err
}
