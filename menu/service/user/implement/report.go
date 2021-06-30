package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/domain"
)

func (impl *implementation) Report(ctx context.Context) (result []domain.CreateStruct, err error) {
	return impl.elasRepo.ReadReport(ctx)
}
