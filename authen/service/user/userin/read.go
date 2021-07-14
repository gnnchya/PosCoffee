package userin

type ReadInput struct {
	ID string `bson:"_id" validate:"required"`
}

