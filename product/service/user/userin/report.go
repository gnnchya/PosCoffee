package userin

import "github.com/gnnchya/PosCoffee/product/domain"

type ReportRange struct{
	From int64
	Until int64
	Format string
}

func (input *ReportRange)ReportInputToUserDomain() (user *domain.ReportValue) {
	return &domain.ReportValue{
		From:  input.From,
		Until: input.Until,
		Format: input.Format,
	}
}

type ReportFilter struct{
	Field string
	Order string
	Format string
}

func (input *ReportFilter)ReportStockInputToUserDomain() (user *domain.ReportOrder) {
	return &domain.ReportOrder{
		Field:  input.Field,
		Order: input.Order,
		Format: input.Format,
	}
}

