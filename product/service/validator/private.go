package validator

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"time"
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

func (v *GoPlayGroundValidator) checkAmount(structLV validator.StructLevel, amount int64, totalAmount int64) {
	if amount != totalAmount{
		structLV.ReportError(amount, "err validation amount", "err validation amount", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkCost(structLV validator.StructLevel, costPerUnit int64, totalAmount int64, totalCost int64) {
	if totalCost != (totalAmount * costPerUnit){
		structLV.ReportError(totalCost, "err validation total cost", "err validation total cost", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkExpDate(structLV validator.StructLevel, exp int64) {
	if exp < time.Now().Unix(){
		structLV.ReportError(exp, "err validation exp date", "err validation exp date", "wrong", "")
	}
}

func (v *GoPlayGroundValidator) checkImportDate(structLV validator.StructLevel, imp int64) {
	if imp > time.Now().Unix(){
		structLV.ReportError(imp, "err validation import date", "err validation import date", "wrong", "")
	}
}

func (v *GoPlayGroundValidator) checkExistVal(structLV validator.StructLevel, val int64) {
	exist, err := v.moneyRepo.CheckExistVal(context.Background(), val)
	if err != nil{
		structLV.ReportError(val, "err validation check exist value", "err validation check exist value", "unique", "")
	}
	if exist{
		structLV.ReportError(val, "err validation check exist value", "err validation check exist value", "unique", "")
	}
}