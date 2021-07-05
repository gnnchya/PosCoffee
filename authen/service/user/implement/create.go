package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) Create(input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "validate error", err
	}
	user := input.CreateInputToUserDomain()

	return user.ID, nil
}
