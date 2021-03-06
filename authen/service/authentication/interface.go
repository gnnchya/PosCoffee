package authentication

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication/authenticationin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/out"
)

//go:generate mockery --name=Service
type Service interface {
	GenerateToken(input *authenticationin.LoginInput) (token *out.Token, err error)
	VerifyToken(accessToken string) (userID *string, err error)
	Logout(input string) (err error)
	RevokeToken(accessToken string) (token bool, err error)
	GetToken(userID, username, password string) (token *out.Token, err error)
}
