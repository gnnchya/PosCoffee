package validator

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/gnnchya/PosCoffee/authorize/domain"
	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

func (v *GoPlayGroundValidator) RolesCreateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	role := structLV.Current().Interface().(rolesin.CreateInput)

	if role.Name.En != "" {
		v.checkRoleNameEN(structLV, "NameEN", role.Name.En)
	}

	if role.Name.Th != "" {
		v.checkRoleNameTH(structLV, "NameTH", role.Name.Th)
	}

	v.checkRolesCreateUniqueTH(ctx, structLV, role.Name.Th)
	v.checkRolesCreateUniqueEN(ctx, structLV, role.Name.En)
	v.checkFormatPermissionCreate(ctx, structLV, role.Permission)
}

func (v *GoPlayGroundValidator) checkRolesCreateUniqueTH(ctx context.Context, structLV validator.StructLevel, nameTh string) {
	filters := []string{
		fmt.Sprintf("name.th:eq:%s", nameTh),
		"deletedAt:isNull",
	}
	if err := v.Repo.Role.Read(ctx, filters, &domain.Roles{}); err == nil {
		structLV.ReportError(nameTh, "NameTh", "NameTh", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkRolesCreateUniqueEN(ctx context.Context, structLV validator.StructLevel, nameEn string) {
	filters := []string{
		fmt.Sprintf("name.en:eq:%s", nameEn),
		"deletedAt:isNull",
	}
	if err := v.Repo.Role.Read(ctx, filters, &domain.Roles{}); err == nil {
		structLV.ReportError(nameEn, "NameEn", "NameEn", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkFormatPermissionCreate(ctx context.Context, structLV validator.StructLevel, permissions []*string) {
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
