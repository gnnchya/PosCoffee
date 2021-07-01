package implement

import (
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"regexp"
)

func MakeFilterEmailORMobileNumber(filter util.Filters, input string) (filters []string) {
	re := regexp.MustCompile("[0-9]{10}")
	ok := re.MatchString(input)
	if !ok {
		return []string{
			filter.MakeFilterEmailString(input),
			"deletedAt:isNull",
		}
	}
	return []string{
		filter.MakeFilterMobileNumberString(input),
		"deletedAt:isNull",
	}
}
