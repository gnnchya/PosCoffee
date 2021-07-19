package userin

type DeleteInput struct {
	ID string `json:"-" validate:"required"`
}
