package validator

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
)

func (v *GoPlayGroundValidator) PermissionsUpdateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	permission := structLV.Current().Interface().(permissionsin.UpdateInput)

	if permission.Method != "" {
		v.checkFormatSpace(structLV, "Method", permission.Method)
	}

	if permission.Endpoint != "" {
		v.checkFormatSpace(structLV, "Endpoint", permission.Endpoint)
	}

	v.checkPermissionsUpdateUnique(ctx, structLV, permission.ID, permission.Method, permission.Endpoint)
}

func (v *GoPlayGroundValidator) checkPermissionsUpdateUnique(ctx context.Context, structLV validator.StructLevel, id string, method string, endpoint string) {
	filters := []string{
		fmt.Sprintf("_id:ne:%s", id),
		fmt.Sprintf("method:eq:%s", method),
		fmt.Sprintf("endpoint:eq:%s", endpoint),
		"deletedAt:isNull",
	}
	if err := v.Repo.Permission.Read(ctx, filters, &domain.Permissions{}); err == nil {
		structLV.ReportError(method, "Method/Endpoint", "Method/Endpoint", "unique", "")
	}
}
