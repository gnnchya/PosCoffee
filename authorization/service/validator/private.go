package validator

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

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

func (v *GoPlayGroundValidator) checkFormatSpace(structLV validator.StructLevel, fieldName string, fieldValue string) {
	re := regexp.MustCompile("^[\\s]+$")
	if re.MatchString(fieldValue) {
		structLV.ReportError(fieldValue, fieldName, fieldName, "match", "")
	}
}

func (v *GoPlayGroundValidator) checkRoleNameTH(structLV validator.StructLevel, fieldName string, fieldValue string) {
	re := regexp.MustCompile("[ก-๛]+")
	ok := re.MatchString(fieldValue)
	if !ok {
		structLV.ReportError(fieldValue, fieldName, fieldName, "match", "")
	}
}

func (v *GoPlayGroundValidator) checkRoleNameEN(structLV validator.StructLevel, fieldName string, fieldValue string) {
	re := regexp.MustCompile("[A-Za-z]+")
	ok := re.MatchString(fieldValue)
	if !ok {
		structLV.ReportError(fieldValue, fieldName, fieldName, "match", "")
	}
}