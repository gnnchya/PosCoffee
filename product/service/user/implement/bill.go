package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/bill"
)

func (impl *implementation) Bill(ctx context.Context, id string) (res []string,err error){
	transaction,err := impl.repo.ReadBill(ctx, id)
	return 	bill.BillInfo(transaction), err
}