package validator

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

func (v *GoPlayGroundValidator) RolesUpdateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	role := structLV.Current().Interface().(rolesin.UpdateInput)

	if role.Name.En != "" {
		v.checkRoleNameEN(structLV, "NameEN", role.Name.En)
	}

	if role.Name.Th != "" {
		v.checkRoleNameTH(structLV, "NameTH", role.Name.Th)
	}

	v.checkRolesUpdateUniqueTH(ctx, structLV, role.ID, role.Name.Th)
	v.checkRolesUpdateUniqueEN(ctx, structLV, role.ID, role.Name.En)
	v.checkFormatPermissionUpdate(ctx, structLV, role.Permission)
}

func (v *GoPlayGroundValidator) checkRolesUpdateUniqueTH(ctx context.Context, structLV validator.StructLevel, id string, nameTh string) {
	filters := []string{
		fmt.Sprintf("_id:ne:%s", id),
		fmt.Sprintf("name.th:eq:%s", nameTh),
		"deletedAt:isNull",
	}
	if err := v.Repo.Role.Read(ctx, filters, &domain.Roles{}); err == nil {
		structLV.ReportError(nameTh, "NameTh", "NameTh", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkRolesUpdateUniqueEN(ctx context.Context, structLV validator.StructLevel, id string, nameEn string) {
	filters := []string{
		fmt.Sprintf("_id:ne:%s", id),
		fmt.Sprintf("name.en:eq:%s", nameEn),
		"deletedAt:isNull",
	}
	if err := v.Repo.Role.Read(ctx, filters, &domain.Roles{}); err == nil {
		structLV.ReportError(nameEn, "NameEn", "NameEn", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkFormatPermissionUpdate(ctx context.Context, structLV validator.StructLevel, permissions []*string) {
	for i := 0; i < len(permissions); i++ {
		filters := []string{
			fmt.Sprintf("_id:eq:%s", *permissions[i]),
			"deletedAt:isNull",
		}
		if err := v.Repo.Permission.Read(ctx, filters, &domain.Permissions{}); err != nil {
			structLV.ReportError(*permissions[i], "Permission", "Permission", "match", "")
		}
	}
}