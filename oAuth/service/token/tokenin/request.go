package tokenin

import (
	"github.com/modern-go/reflect2"

	"github.com/gnnchya/PosCoffee/oAuth/domain"
)

type RequestInput struct {
	ID               string `json:"id"`
	UserId           string `json:"user_id" form:"user_id"`
	ClientID         string `json:"client_id" form:"client_id" binding:"required"`
	ClientSecret     string `json:"client_secret" form:"client_secret" binding:"required"`
	GrantType        string `json:"grant_type" form:"grant_type" binding:"required"`
	Scope            string `json:"scope" form:"scope"`
	RedirectUri      string `json:"redirect_uri" form:"redirect_uri"`
	OauthProvider    string `json:"oauth_provider" form:"oauth_provider"`
	Expired          int64  `json:"expired"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiryDate       int64  `json:"expiry_date"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
}

func (input *RequestInput) ToDomain() (token *domain.OauthToken) {
	if reflect2.IsNil(input) {
		return &domain.OauthToken{}
	}

	return &domain.OauthToken{
		ID:               input.ID,
		UserId:           input.UserId,
		Expired:          input.Expired,
		OauthProvider:    input.OauthProvider,
		AccessToken:      input.AccessToken,
		RefreshToken:     input.RefreshToken,
		ExpiryDate:       input.ExpiryDate,
		RefreshExpiresIn: input.RefreshExpiresIn,
	}
}