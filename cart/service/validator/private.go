package validator

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/cart/domain"
	"github.com/gnnchya/PosCoffee/cart/service/user/userin"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
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


func (v *GoPlayGroundValidator) checkAmountStruct(structLV validator.StructLevel, menu userin.Input){
	if len(menu.Menu) < 0{
		structLV.ReportError(len(menu.Menu), "err validation for menu", "err validation for menu", "zero", "")
	}
}

func (v *GoPlayGroundValidator) checkUpdateStruct(structLV validator.StructLevel, cart domain.CreateStruct){
	if len(cart.Menu) == 0{
		structLV.ReportError(cart, "err validation cart is 0", "err validation cart is 0", "zero", "")
	}
	for i, j := range cart.Menu{
		if j.Amount <= 0{
			structLV.ReportError(cart, "err validation for cart #"+strconv.Itoa(i), "err validation for cart #"+strconv.Itoa(i), "zero", "")
		}
	}
}