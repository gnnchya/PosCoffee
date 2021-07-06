package authentication

import (
	"github.com/gnnchya/PosCoffee/authen/service/authentication/autenitcationin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/out"
)

//go:generate mockery --name=Service
type Service interface {
	GenerateToken(input *authenitcationin.LoginInput) (token *out.Token, err error)
	VerifyToken(accessToken string) (userID *string, err error)
	Logout(input string) (err error)
}
