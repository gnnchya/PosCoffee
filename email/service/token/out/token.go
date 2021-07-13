package out

import "github.com/gnnchya/PosCoffee/email/domain"

type TokenView struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expired      int64  `json:"expired"`
}

func TokenToView(token *domain.TokenStruct) (view *TokenView) {
	return &TokenView{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expired:      token.AccessExpire,
	}
}
