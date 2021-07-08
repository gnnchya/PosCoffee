package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) checkConsumerExistID(structLV validator.StructLevel, ID string) {
	exist, _ := v.consumerRepo.CheckExistID(context.Background(), ID)
	if exist{
		structLV.ReportError(ID, "err2 validation check exist ID", "err validation check exist ID", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkTokenExistID(structLV validator.StructLevel, ID string) {
	exist, _ := v.tokenRepo.CheckExistID(context.Background(), ID)
	if exist{
		structLV.ReportError(ID, "err2 validation check exist ID", "err validation check exist ID", "unique", "")
	}
}