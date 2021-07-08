package rolesin

type ReadInput struct {
	ID string `bson:"-" validate:"required"`
} // @Name RolesReadInput
