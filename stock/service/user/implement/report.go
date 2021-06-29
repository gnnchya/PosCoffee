package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

func (impl *implementation) Report(ctx context.Context) (result []domain.CreateStruct, err error) {
	result, err = impl.repo.Report(ctx)
	if err != nil {
		fmt.Println("err in report", err)
		return result, err
	}
	return result, err
}