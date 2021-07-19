package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
)

func (impl *implementation) Create(ctx context.Context, input *consumerin.CreateConsumerInput) (View *consumerin.ConsumerOutput, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return View, err
	}

	clientId, err := impl.uuid4.NewRandom()
	if err != nil {
		return View, err
	}
	clientSecret, err := impl.uuid4.NewRandom()
	if err != nil {
		return View, err
	}

	input.ID = impl.uuid.Generate()
	input.ClientID = clientId
	input.ClientSecret = clientSecret

	consumer := input.ToDomain()

	err = impl.repo.Create(ctx, consumer)
	if err != nil {
		return View, err
	}
	return input.View(), nil
}
