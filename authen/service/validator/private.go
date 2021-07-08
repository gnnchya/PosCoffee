package validator

import (
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

func (v *GoPlayGroundValidator) checkValidUsername(structLV validator.StructLevel, name string) {
	re := regexp.MustCompile("^[A-Za-z]\\w{5, 29}$")
	ok := re.MatchString(name)
	if !ok {
		structLV.ReportError(name, "username", "username", "not valid", "")
	}
}