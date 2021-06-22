package validator

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/gnnchya/PosCoffee/menu/domain"

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

func (v *GoPlayGroundValidator) checkNameUnique(ctx context.Context, structLV validator.StructLevel, name string) (user *domain.InsertQ) {
	a, err := v.userRepo.CheckExistName(ctx, name)
	if err != nil {
		structLV.ReportError(err, "err validation", "err validation", "error from database", "")
	}
	if a == true {
		structLV.ReportError(name, "name", "name", "unique", "")
	}
	return user
}

//func (v *GoPlayGroundValidator) checkUserActualNameUnique(ctx context.Context, structLV validator.StructLevel, name string) (user *domain.InsertQ) {
//	a, _ := v.userRepo.CheckExistActualName(ctx, name)
//	// if err != nil {
//	// 	structLV.ReportError(err, "err validation", "err validation", "error from database", "")
//	// }
//	if a == true {
//		structLV.ReportError(name, "actual_name", "actual_name", "unique", "")
//	}
//	return user
//}

func (v *GoPlayGroundValidator) checkNameUniqueUpdate(ctx context.Context, structLV validator.StructLevel, name string, id string) (user *domain.UpdateQ) {
	n, err := v.userRepo.CheckExistName(ctx, name)
	if n == true { //jer name
		if an == true { //
			temp, _ := v.userRepo.View(ctx, id)
			if temp.Name != name {
				structLV.ReportError(actualName, "actual_name", "actual_name", "unique", "")
			}
		}
	}
	return user
}
