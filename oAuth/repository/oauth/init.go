package oauth

import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/store"
)

type Oauth struct {
	manager *manage.Manager
}

func New() (oauth2 *Oauth) {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	oauth2 = &Oauth{
		manager,
	}
	return oauth2
}
