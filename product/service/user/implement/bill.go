package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/bill"
)

func (impl *implementation) Bill(ctx context.Context, id string, paid int64)[]string{
	res,_ := impl.repo.ReadBill(ctx, id)
	result := bill.BillInfo(res,paid)
	return result
}