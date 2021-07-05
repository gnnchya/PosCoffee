package implement

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) Create(input *userin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		fmt.Println("validate", err)
		return "validate error", err
	}
	user := input.CreateInputToUserDomain()
	fmt.Println("user input create:", user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
