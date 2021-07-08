package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *rolesin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	input.ID = impl.uuid.Generate()
	role := input.ToDomain()

	_, err = impl.repo.Create(ctx, role)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return role.ID, nil
}
