package consumerin

import (
	"github.com/gnnchya/PosCoffee/oAuth/domain"
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"
)

type CreateConsumerInput struct {
	ID              string `json:"id"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	Scope           string `json:"scope"`
}

func (input *CreateConsumerInput) ToDomain() (consumer *domain.ConsumerStruct) {
	if reflect2.IsNil(input) {
		return &domain.ConsumerStruct{}
	}

	return &domain.ConsumerStruct{
		ID:              input.ID,
		ClientID:        input.ClientID,
		ClientSecret:    input.ClientSecret,
		Scope:           input.Scope,
		CreatedAt:       carbon.Now().Unix(),
	}
}
