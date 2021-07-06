package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
	"github.com/gnnchya/PosCoffee/oAuth/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *consumerin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	clientId, err := impl.uuid4.NewRandom()
	if err != nil {
		return "", util.RepoCreateErr(err)
	}
	clientSecret, err := impl.uuid4.NewRandom()
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	input.ID = impl.uuid.Generate()
	input.ClientId = clientId
	input.ClientSecret = clientSecret

	consumer := input.ToDomain()

	_, err = impl.repo.Create(ctx, consumer)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return consumer.ID, nil
}
