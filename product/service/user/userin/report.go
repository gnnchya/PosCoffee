package userin

import "github.com/gnnchya/PosCoffee/product/domain"

type ReportRange struct{
	From int64
	Until int64
}

func (input *ReportRange)ReportInputToUserDomain() (user *domain.ReportValue) {
	return &domain.ReportValue{
		From:  input.From,
		Until: input.Until,
	}
}

