package validator

type Validator interface {
	Validate(item interface{}) (err error)
}
