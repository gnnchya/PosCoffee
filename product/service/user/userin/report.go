package userin

import "github.com/gnnchya/PosCoffee/product/domain"

type ReportRange struct{
	From 	int64	`bson:"from" json:"from"`
	Until 	int64	`bson:"until" json:"until"`
	Format 	string	`bson:"format" json:"format"`
}

func (input *ReportRange)ReportInputToUserDomain() (user *domain.ReportValue) {
	return &domain.ReportValue{
		From:  input.From,
		Until: input.Until,
		Format: input.Format,
	}
}

type ReportFilter struct{
	Field 	string	`bson:"field" json:"field"`
	Order 	string	`bson:"order" json:"order"`
	Format 	string	`bson:"format" json:"format"`
}

func (input *ReportFilter)ReportStockInputToUserDomain() (user *domain.ReportOrder) {
	return &domain.ReportOrder{
		Field:  input.Field,
		Order: input.Order,
		Format: input.Format,
	}
}