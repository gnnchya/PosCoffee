package permissionsin

type ReadInput struct {
	ID string `bson:"_id" json:"id"validate:"required"`
}
