package oauth

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
)

func (manage Oauth) ClientStore(clientID, clientSecret, domain string) oauth2.ClientStore {
	clientStore := store.NewClientStore()
	err := clientStore.Set(clientID, &models.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: domain,
	})

	if err != nil {
		return nil
	}

	manage.manager.MapClientStorage(clientStore)

	return clientStore
}
