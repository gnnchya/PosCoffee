package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"strings"
)

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

func (v *GoPlayGroundValidator) pageOptionFilterValidation(structLV validator.StructLevel, filter string) {
	fragments := strings.Split(filter, ":")
	if len(fragments) < 2 {
		structLV.ReportError(filter, "filters", "filters", "length", "")
		return
	}

	support := map[string]bool{
		"ne":        true,
		"like":      true,
		"eq":        true,
		"eqInt":     true,
		"isNull":    true,
		"isNotNull": true,
		"id":        true,
	}
	if _, ok := support[fragments[1]]; !ok {
		structLV.ReportError(filter, "filters", "filters", "support", "")
		return
	}

	switch fragments[1] {
	case "isNull":
		if len(fragments) == 2 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	case "isNotNull":
		if len(fragments) == 2 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	default:
		if len(fragments) == 3 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	}
}

func (v *GoPlayGroundValidator) pageOptionSortValidation(structLV validator.StructLevel, sort string) {
	fragments := strings.Split(sort, ":")
	if len(fragments) < 2 {
		structLV.ReportError(sort, "sorts", "sorts", "length", "")
		return
	}

	support := map[string]bool{
		"asc":  true,
		"desc": true,
	}
	if _, ok := support[fragments[1]]; !ok {
		structLV.ReportError(sort, "sorts", "sorts", "support", "")
		return
	}
}
