package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func (v *GoPlayGroundValidator) checkName(structLV validator.StructLevel, name string) {
	if name != "Proud" {
		structLV.ReportError("not Proud", "name", "name", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkTH(structLV validator.StructLevel, name string) {
	re := regexp.MustCompile("[A-Za-z]+")
	ok := re.MatchString(name)
	if !ok {
		structLV.ReportError(name, "err validation", "err validation", "match", "")
	}
}

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