package tokenin

import (
	"github.com/modern-go/reflect2"

	"github.com/gnnchya/PosCoffee/oAuth/domain"
)

type RequestInput struct {
	ID               string `json:"id"`
	UID          	 string `json:"uid" form:"uid" binding:"required"`
	ClientID         string `json:"client_id" form:"client_id" binding:"required"`
	ClientSecret     string `json:"client_secret" form:"client_secret" binding:"required"`
	GrantType        string `json:"grant_type" form:"grant_type" binding:"required"`
	Scope            string `json:"scope" form:"scope"`
	RedirectUri      string `json:"redirect_uri" form:"redirect_uri"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	AccessExpire     int64  `json:"access_expire"`
	RefreshExpire 	 int64  `json:"refresh_expires_in"`
	CreatedAt 		 int64	`bson:"created_at" json:"created_at"`
}

func (input *RequestInput) ToDomain() (token *domain.TokenStruct) {
	if reflect2.IsNil(input) {
		return &domain.TokenStruct{}
	}

	return &domain.TokenStruct{
		ID:               input.ID,
		UID:              input.UID,
		AccessToken:      input.AccessToken,
		RefreshToken:     input.RefreshToken,
		AccessExpire:     input.AccessExpire,
		RefreshExpire: 	  input.RefreshExpire,
		CreatedAt: 		  input.CreatedAt,
	}
}