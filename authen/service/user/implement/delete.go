package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
)

func (impl *implementation) Delete(ctx context.Context, input *userin.DeleteInput) (ID string, err error) {
	fmt.Println("user input delete:", input)
	if err != nil {
		return "", err
	}
	return input.ID, err
}


