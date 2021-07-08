package out

import "github.com/gnnchya/PosCoffee/oAuth/domain"

type TokenView struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessExpire int64  `json:"expired"`
}

func TokenToView(token *domain.TokenStruct) (view *TokenView) {
	return &TokenView{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		AccessExpire: token.AccessExpire,
	}
}
