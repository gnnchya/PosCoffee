package validator

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"regexp"
	"strconv"

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


func (v *GoPlayGroundValidator) checkCreateStruct(structLV validator.StructLevel, cart []domain.Cart){
	if len(cart) == 0{
		structLV.ReportError(cart, "err validation cart is 0", "err validation cart is 0", "zero", "")
	}
	for i, j := range cart{
		if j.Amount <= 0{
			structLV.ReportError(cart, "err validation for cart #"+strconv.Itoa(i), "err validation for cart #"+strconv.Itoa(i), "zero", "")
		}
	}
}