package out

import "github.com/gnnchya/PosCoffee/oAuth/domain"

type ValidateTokenView struct {
	UserId  string `json:"user_id"`
	Expired int64  `json:"expired"`
}

func ValidateTokenToView(token *domain.TokenStruct) (view *ValidateTokenView) {
	return &ValidateTokenView{
		UserId:  token.UID,
		Expired: token.CreatedAt,
	}
}
