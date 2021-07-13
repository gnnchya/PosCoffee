package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/email/service/consumer/consumerin"
)

func (impl *implementation) Create(ctx context.Context, input *consumerin.CreateConsumerInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", err
	}

	clientId, err := impl.uuid4.NewRandom()
	if err != nil {
		return "", err
	}
	clientSecret, err := impl.uuid4.NewRandom()
	if err != nil {
		return "", err
	}

	input.ID = impl.uuid.Generate()
	input.ClientID = clientId
	input.ClientSecret = clientSecret

	consumer := input.ToDomain()

	err = impl.repo.Create(ctx, consumer)
	if err != nil {
		return "", err
	}
	//fmt.Println("clientID", consumer.ClientID)
	//fmt.Println("clientSecret", consumer.ClientSecret)
	return consumer.ID, nil
}
