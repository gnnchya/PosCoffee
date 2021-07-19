package out

type RolesView struct {
	ID          string        `json:"id"`
	Name        *Lang         `json:"name"`
	Permissions []*Permission `json:"permissions"`
	CreatedAt   int64         `json:"created_at"`
	UpdatedAt   int64         `json:"updated_at"`
} // @Name RolesView

type Lang struct {
	Th string `json:"th"`
	En string `json:"en"`
}// @Name RolesLangView

type Permission struct {
	ID       string `json:"id"`
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

func RolesToView(roles *RolesView) (view *RolesView) {
	view = &RolesView{
		ID:          roles.ID,
		Name:        &Lang{},
		Permissions: []*Permission{},
		CreatedAt:   roles.CreatedAt,
		UpdatedAt:   roles.UpdatedAt,
	}

	view.setName(roles.Name)
	view.setPermissions(roles.Permissions)

	return view
}

func (view *RolesView) setName(name *Lang) {
	if name != nil {
		view.Name = &Lang{
			Th: name.Th,
			En: name.En,
		}
	}
}

func (view *RolesView) setPermissions(permissions []*Permission) {
	if len(permissions) > 0 {
		view.Permissions = make([]*Permission, len(permissions))
		for i, permission := range permissions {
			view.Permissions[i] = &Permission{
				ID:       permission.ID,
				Method:   permission.Method,
				Endpoint: permission.Endpoint,
			}
		}
	}
}