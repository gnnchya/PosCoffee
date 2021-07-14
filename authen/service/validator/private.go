package validator

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)


func (v *GoPlayGroundValidator) checkTH(structLV validator.StructLevel, name string) {
	re := regexp.MustCompile("[A-Za-z]+")
	ok := re.MatchString(name)
	if !ok {
		structLV.ReportError(name, "err validation", "err validation", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkUsername(structLV validator.StructLevel, name string) {
	filters := []string{fmt.Sprintf("username:eq:%s", name)}
	if found , err := v.repo.Count(context.Background(), filters); err != nil{
		structLV.ReportError(err, "username", "username", "error", "")
	} else if found >0{
		structLV.ReportError(name, "username", "username", "unique", "")
	}
	re := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]{7,19}$")
	ok := re.MatchString(name)
	if !ok {
		structLV.ReportError(name, "username", "username", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkFirstNameTH(structLV validator.StructLevel, firstNameTH string) {
	re := regexp.MustCompile("^[ก-๛]+$")
	ok := re.MatchString(firstNameTH)
	if !ok {
		structLV.ReportError(firstNameTH, "firstNameTH", "firstNameTH", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkLastNameTH(structLV validator.StructLevel, lastNameTH string) {
	re := regexp.MustCompile("^[ก-๛]+$")
	ok := re.MatchString(lastNameTH)
	if !ok {
		structLV.ReportError(lastNameTH, "lastNameTH", "lastNameTH", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkFirstNameEN(structLV validator.StructLevel, firstNameEN string) {
	re := regexp.MustCompile("^[A-Za-z]+$")
	ok := re.MatchString(firstNameEN)
	if !ok {
		structLV.ReportError(firstNameEN, "firstNameEN", "firstNameEN", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkLastNameEN(structLV validator.StructLevel, lastNameEN string) {
	re := regexp.MustCompile("^[A-Za-z]+$")
	ok := re.MatchString(lastNameEN)
	if !ok {
		structLV.ReportError(lastNameEN, "lastNameEN", "lastNameEN", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkPrefixEN(structLV validator.StructLevel, prefixEN string) {
	re := regexp.MustCompile("^[a-zA-z.]{2,}$")
	ok := re.MatchString(prefixEN)
	if !ok {
		structLV.ReportError(prefixEN, "prefixEN", "prefixEN", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkPrefixTH(structLV validator.StructLevel, prefixTH string) {
	re := regexp.MustCompile("^[ก-๛.]{2,}$")
	ok := re.MatchString(prefixTH)
	if !ok {
		structLV.ReportError(prefixTH, "prefixTH", "prefixTH", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkFormatBirthOfDate(structLV validator.StructLevel, birthDate int) {
	re := regexp.MustCompile("[0-9]+")
	birthDateString := strconv.Itoa(birthDate)
	ok := re.MatchString(birthDateString)
	if !ok {
		structLV.ReportError(birthDate, "birthDate", "birthDate", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkFormatGender(structLV validator.StructLevel, gender string) {
	genders := []string{
		"Male",
		"Female",
		"Others",
	}
	var ok bool
	for _, v := range genders {
		if gender == v {
			ok = true
		}
	}
	if !ok {
		structLV.ReportError(gender, "gender", "gender", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkBankAccount(structLV validator.StructLevel, bank string){
	re := regexp.MustCompile("^[0-9]*$")
	ok := re.MatchString(bank)
	if !ok {
		structLV.ReportError(bank, "bankAccount", "bankAccount", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkMobileNumberUnique(ctx context.Context, structLV validator.StructLevel, mobileNumber string) {
	filters := []string{fmt.Sprintf("mobile_number:eq:%s", mobileNumber)}
	if found , err := v.repo.Count(ctx, filters); err != nil{
		structLV.ReportError(err, "mobile_number", "mobile_number", "error", "")
	} else if found >0{
		structLV.ReportError(mobileNumber, "mobile_number", "mobile_number", "unique", "")
	}
	re := regexp.MustCompile("^[0-9]*$")
	ok := re.MatchString(mobileNumber)
	if !ok {
		structLV.ReportError(mobileNumber, "mobile number", "mobile number", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkFormatEmail(structLV validator.StructLevel, email string) {
	filters := []string{fmt.Sprintf("meta_data.email:eq:%s", email)}
	if found , err := v.repo.Count(context.Background(), filters); err != nil{
		structLV.ReportError(err, "email", "email", "error", "")
	} else if found > 0{
		structLV.ReportError(email, "email", "email", "unique", "")
	}
	
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	ok := re.MatchString(email)
	if !ok {
		structLV.ReportError(email, "email", "email", "match", "")
	}
}

