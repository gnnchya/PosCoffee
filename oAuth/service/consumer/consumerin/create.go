package consumerin

import (
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"
)

type CreateInput struct {
	ID              string `json:"id" swaggerignore:"true"`
	ApplicationName string `json:"application_name" validate:"required"`
	ClientId        string `json:"client_id" swaggerignore:"true"`
	ClientSecret    string `json:"client_secret" swaggerignore:"true"`
	RedirectUri     string `json:"redirect_uri"`
	Scope           string `json:"scope"`
} // @Name ConsumerCreateInput

func (input *CreateInput) ToDomain() (consumer *domain.Consumer) {
	if reflect2.IsNil(input) {
		return &domain.Consumer{}
	}

	return &domain.Consumer{
		ID:              input.ID,
		ApplicationName: input.ApplicationName,
		ClientId:        input.ClientId,
		ClientSecret:    input.ClientSecret,
		RedirectUri:     input.RedirectUri,
		Scope:           input.Scope,
		CreatedAt:       carbon.Now().Unix(),
	}
}
