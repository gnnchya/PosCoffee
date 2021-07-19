package validator

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserUpdateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(userin.UpdateInput)
	v.checkBankAccount(structLV, input.MetaData.BankAccount)
	v.checkPrefixEN(structLV, input.MetaData.Prefix.EN)
	v.checkPrefixTH(structLV, input.MetaData.Prefix.TH)
	v.checkFirstNameEN(structLV, input.MetaData.Name.EN)
	v.checkFirstNameTH(structLV, input.MetaData.Name.TH)
	v.checkLastNameEN(structLV, input.MetaData.Lastname.EN)
	v.checkLastNameTH(structLV, input.MetaData.Lastname.TH)
	v.checkFormatEmail(structLV, input.MetaData.Email)
	v.checkFormatBirthOfDate(structLV, int(input.MetaData.BirthDate))
	v.checkFormatGender(structLV, input.MetaData.Gender)
	v.checkMobileNumberUnique(ctx, structLV, input.MetaData.MobileNumber)
}

