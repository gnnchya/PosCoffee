package validator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) checkName(structLV validator.StructLevel, name string) {
	fmt.Println("here")
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

func (v *GoPlayGroundValidator) checkNameUnique(ctx context.Context, structLV validator.StructLevel, name string) {
	a, err := v.elasRepo.CheckExistName(ctx, name)
	if err != nil {
		structLV.ReportError(err, "err validation", "err validation", "error from database", "")
	}
	if a == true {
		structLV.ReportError(name, "name", "name", "unique", "")
	}
}


func (v *GoPlayGroundValidator) checkNameUniqueUpdate(ctx context.Context, structLV validator.StructLevel, name string, id string)  {
	n, err := v.elasRepo.CheckExistName(ctx, name)
	if err != nil{
		structLV.ReportError(err, "err update validation", "err update validation", "err update validation", "")
	}
	if n == true { //jer name

		temp, _ := v.elasRepo.Read(id, ctx)
		if temp.ID != id {
			structLV.ReportError(name, "name", "name", "unique", "")
		}

	}
}
