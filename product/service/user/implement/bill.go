package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/service/bill"
	"github.com/jung-kurt/gofpdf"
)

func (impl *implementation) Bill(ctx context.Context, id string) (pdf *gofpdf.Fpdf,err error){
	transaction,err := impl.repo.ReadBill(ctx, id)
	if err != nil {
		return pdf, err
	}
	order := bill.BillInfo(transaction)
	return 	bill.GeneratePdf("bill.pdf", order)
}