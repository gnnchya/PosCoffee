package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*domain.UserStruct, err error) {
	total, records, err := impl.repo.List(ctx, opt, &domain.UserStruct{})
	if err != nil || len(records) == 0 {
		return 0, nil, err
	}

	items = make([]*domain.UserStruct, len(records))
	for i, record := range records {
		items[i] = record.(*domain.UserStruct)
	}
	return total, items, nil
}
