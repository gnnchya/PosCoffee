package implement

import (
	"fmt"
)

func makePermissionIDFilters(permissionID string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", permissionID),
		"deletedAt:isNull",
	}
}

func makeDeletedAtIsNullFilterString() (filter string) {
	return "deletedAt:isNull"
}
