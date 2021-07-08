package permissionsin

type ReadInput struct {
	ID string `bson:"-" validate:"required"`
} // @Name PermissionsReadInput
