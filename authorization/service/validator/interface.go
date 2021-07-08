package validator

//go:generate mockery --name=Validator
type Validator interface {
	Validate(item interface{}) (err error)
}
