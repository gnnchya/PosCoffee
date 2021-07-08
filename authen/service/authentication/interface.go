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
}
