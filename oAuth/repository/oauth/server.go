package oauth

import "github.com/go-oauth2/oauth2/v4/server"

func (manage Oauth) NewServer() *server.Server {
	return server.NewDefaultServer(manage.manager)
}
