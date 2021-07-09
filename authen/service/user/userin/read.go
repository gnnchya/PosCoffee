package userin

type ReadInput struct {
	ID string `bson:"-" validate:"required"`
}

