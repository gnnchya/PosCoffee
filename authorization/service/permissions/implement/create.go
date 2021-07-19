package implement

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
	"github.com/gnnchya/PosCoffee/authorize/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *permissionsin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	input.ID = impl.uuid.Generate()
	permission := input.ToDomain()

	_, err = impl.repo.Create(ctx, permission)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return permission.ID, nil
}
