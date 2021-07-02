package implement

import (
	"context"
)

func (impl *implementation) Bill(ctx context.Context, id string){
	_,_ = impl.repo.ReadBill(ctx, id)
	return
}