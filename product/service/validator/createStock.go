package validator

import (
	"github.com/gnnchya/PosCoffee/product/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserCreateStockStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(userin.CreateStockInput)

	v.checkAmount(structLV, input.Amount, input.TotalAmount)
	v.checkCost(structLV, input.CostPerUnit, input.TotalAmount, input.TotalCost)
	v.checkExpDate(structLV, input.EXPDate)
	v.checkImportDate(structLV, input.ImportDate)
}