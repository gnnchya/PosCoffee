package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"reflect"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (err error) {
	var (
		hashPassword []byte
	)
	err = impl.validator.Validate(input)
	if err != nil {
		return err
	}

	filters := impl.filter.MakeIdFilters(input.ID)
	users := &domain.UserStruct{}
	err = impl.repo.Read(ctx, filters, users)
	if err != nil {
		return err
	}

	val := reflect.ValueOf(hashPassword)
	if val.IsZero() {
		hashPassword = []byte(users.Password)
	}

	updateInput := userin.UpdateInput{
		ID:             input.ID,
		UID:			input.UID,
		Username: 		input.Username,
		Password:       input.Password,
		MetaData:       input.MetaData,
		RoleID:         input.RoleID,
	}
	update := updateInput.ToDomain()
	update.CreatedAt = users.CreatedAt
	update.ID = ""
	err = impl.repo.Update(ctx, filters, update)
	if err != nil {
		return err
	}

	return nil
}
